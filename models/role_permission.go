package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	DisplayName string         `json:"display_name" gorm:"not null;size:150" binding:"required"`
	Description string         `json:"description" gorm:"size:500"`
	Module      string         `json:"module" gorm:"not null;size:50;index" binding:"required"`
	Action      string         `json:"action" gorm:"not null;size:50" binding:"required"`
	Resource    string         `json:"resource" gorm:"size:50"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Roles []Role `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	DisplayName string         `json:"display_name" gorm:"not null;size:150" binding:"required"`
	Description string         `json:"description" gorm:"size:500" binding:"required"`
	Level       int            `json:"level" gorm:"not null;default:1;check:level >= 0 AND level <= 5" binding:"required,min=0,max=5"`
	Color       string         `json:"color" gorm:"size:20;default:'blue'" binding:"required"`
	IsDefault   bool           `json:"is_default" gorm:"default:false"`
	IsSystem    bool           `json:"is_system" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	UserCount   int            `json:"user_count" gorm:"default:0"`
	CreatedBy   uint           `json:"created_by" gorm:"index"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
	Users       []User       `json:"users,omitempty" gorm:"foreignKey:RoleID"`
	Creator     *User        `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	RoleID       uint      `json:"role_id" gorm:"primaryKey;autoIncrement"`
	PermissionID uint      `json:"permission_id" gorm:"primaryKey"`
	GrantedBy    uint      `json:"granted_by" gorm:"index"`
	GrantedAt    time.Time `json:"granted_at" gorm:"autoCreateTime"`

	Role          Role       `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	Permission    Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
	GrantedByUser *User      `json:"granted_by_user,omitempty" gorm:"foreignKey:GrantedBy"`
}

type PermissionModule struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null;size:50" binding:"required"`
	DisplayName string    `json:"display_name" gorm:"not null;size:100" binding:"required"`
	Description string    `json:"description" gorm:"size:300"`
	Icon        string    `json:"icon" gorm:"size:50"`
	SortOrder   int       `json:"sort_order" gorm:"default:0"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Permissions []Permission `json:"permissions,omitempty" gorm:"foreignKey:Module;references:Name"`
}

type PermissionTemplate struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	Description string       `json:"description" gorm:"size:500"`
	Category    string       `json:"category" gorm:"size:50;index"`
	IsActive    bool         `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:template_permissions;"`
}

type TemplatePermission struct {
	TemplateID   uint `json:"template_id" gorm:"primaryKey"`
	PermissionID uint `json:"permission_id" gorm:"primaryKey"`

	Template   PermissionTemplate `json:"template,omitempty" gorm:"foreignKey:TemplateID"`
	Permission Permission         `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
}

type RoleHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	RoleID      uint      `json:"role_id" gorm:"not null;index"`
	Action      string    `json:"action" gorm:"not null;size:50"` // CREATE, UPDATE, DELETE, PERMISSION_GRANT, PERMISSION_REVOKE
	FieldName   string    `json:"field_name" gorm:"size:50"`
	OldValue    string    `json:"old_value" gorm:"type:text"`
	NewValue    string    `json:"new_value" gorm:"type:text"`
	ChangedByID uint      `json:"changed_by" gorm:"not null;index"`
	Reason      string    `json:"reason" gorm:"size:500"`
	IPAddress   string    `json:"ip_address" gorm:"size:45"`
	UserAgent   string    `json:"user_agent" gorm:"size:500"`
	CreatedAt   time.Time `json:"created_at"`

	Role      Role `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	ChangedBy User `json:"changed_by_user,omitempty" gorm:"foreignKey:ChangedBy"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.IsDefault {
		tx.Model(&Role{}).Where("is_default = ?", true).Update("is_default", false)
	}
	return nil
}

func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	if r.IsDefault {
		tx.Model(&Role{}).Where("id != ? AND is_default = ?", r.ID, true).Update("is_default", false)
	}
	return nil
}

func (r *Role) BeforeDelete(tx *gorm.DB) error {
	if r.IsSystem {
		return gorm.ErrRecordNotFound // Or custom error
	}

	var userCount int64
	tx.Model(&User{}).Where("role_id = ?", r.ID).Count(&userCount)
	if userCount > 0 {
		return gorm.ErrRecordNotFound // Or custom error: cannot delete role with users
	}

	return nil
}

func (r *Role) UpdateUserCount(tx *gorm.DB) error {
	var count int64
	tx.Model(&User{}).Where("role_id = ?", r.ID).Count(&count)
	return tx.Model(r).Update("user_count", count).Error
}

func (r *Role) HasPermission(tx *gorm.DB, permissionName string) bool {
	var count int64
	tx.Table("role_permissions").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ? AND permissions.name = ? AND permissions.is_active = ?",
			r.ID, permissionName, true).
		Count(&count)
	return count > 0
}

func (r *Role) GetPermissionsByModule(tx *gorm.DB) (map[string][]Permission, error) {
	var permissions []Permission
	err := tx.Model(r).Association("Permissions").Find(&permissions)
	if err != nil {
		return nil, err
	}

	modulePermissions := make(map[string][]Permission)
	for _, permission := range permissions {
		if permission.IsActive {
			modulePermissions[permission.Module] = append(modulePermissions[permission.Module], permission)
		}
	}

	return modulePermissions, nil
}

func (p *Permission) IsValidAction() bool {
	validActions := []string{"view", "create", "edit", "delete", "manage", "export", "import", "approve"}
	for _, action := range validActions {
		if p.Action == action {
			return true
		}
	}
	return false
}

func (p *Permission) GetFullPermissionName() string {
	if p.Resource != "" {
		return p.Module + ":" + p.Action + ":" + p.Resource
	}
	return p.Module + ":" + p.Action
}

func GetDefaultPermissions() []Permission {
	return []Permission{
		// Dashboard permissions
		{Name: "dashboard.view", DisplayName: "View Dashboard", Module: "dashboard", Action: "view", Description: "Access to main dashboard"},
		{Name: "dashboard.analytics", DisplayName: "View Analytics", Module: "dashboard", Action: "analytics", Description: "Access to analytics and reports"},

		// Product permissions
		{Name: "products.view", DisplayName: "View Products", Module: "products", Action: "view", Description: "View product listings"},
		{Name: "products.create", DisplayName: "Create Products", Module: "products", Action: "create", Description: "Add new products"},
		{Name: "products.edit", DisplayName: "Edit Products", Module: "products", Action: "edit", Description: "Modify existing products"},
		{Name: "products.delete", DisplayName: "Delete Products", Module: "products", Action: "delete", Description: "Remove products"},
		{Name: "products.manage_inventory", DisplayName: "Manage Inventory", Module: "products", Action: "manage", Resource: "inventory", Description: "Manage stock levels"},

		// Order permissions
		{Name: "orders.view", DisplayName: "View Orders", Module: "orders", Action: "view", Description: "View order listings"},
		{Name: "orders.create", DisplayName: "Create Orders", Module: "orders", Action: "create", Description: "Create new orders"},
		{Name: "orders.edit", DisplayName: "Edit Orders", Module: "orders", Action: "edit", Description: "Modify existing orders"},
		{Name: "orders.delete", DisplayName: "Delete Orders", Module: "orders", Action: "delete", Description: "Remove orders"},
		{Name: "orders.process", DisplayName: "Process Orders", Module: "orders", Action: "manage", Resource: "processing", Description: "Process and fulfill orders"},

		// Customer permissions
		{Name: "customers.view", DisplayName: "View Customers", Module: "customers", Action: "view", Description: "View customer listings"},
		{Name: "customers.create", DisplayName: "Create Customers", Module: "customers", Action: "create", Description: "Add new customers"},
		{Name: "customers.edit", DisplayName: "Edit Customers", Module: "customers", Action: "edit", Description: "Modify customer information"},
		{Name: "customers.delete", DisplayName: "Delete Customers", Module: "customers", Action: "delete", Description: "Remove customers"},

		// User management permissions
		{Name: "users.view", DisplayName: "View Users", Module: "users", Action: "view", Description: "View system users"},
		{Name: "users.create", DisplayName: "Create Users", Module: "users", Action: "create", Description: "Add new users"},
		{Name: "users.edit", DisplayName: "Edit Users", Module: "users", Action: "edit", Description: "Modify user information"},
		{Name: "users.delete", DisplayName: "Delete Users", Module: "users", Action: "delete", Description: "Remove users"},
		{Name: "roles.manage", DisplayName: "Manage Roles", Module: "users", Action: "manage", Resource: "roles", Description: "Manage user roles and permissions"},

		// Reports permissions
		{Name: "reports.view", DisplayName: "View Reports", Module: "reports", Action: "view", Description: "Access to reports"},
		{Name: "reports.export", DisplayName: "Export Reports", Module: "reports", Action: "export", Description: "Export report data"},

		// Settings permissions
		{Name: "settings.view", DisplayName: "View Settings", Module: "settings", Action: "view", Description: "View system settings"},
		{Name: "settings.edit", DisplayName: "Edit Settings", Module: "settings", Action: "edit", Description: "Modify system settings"},
	}
}

// GetDefaultRoles returns the default system roles
func GetDefaultRoles() []Role {
	return []Role{
		{
			Name:        "super_admin",
			DisplayName: "Super Administrator",
			Description: "Full system access with all permissions",
			Level:       5,
			Color:       "red",
			IsSystem:    true,
			IsDefault:   false,
			IsActive:    true,
		},
		{
			Name:        "admin",
			DisplayName: "Administrator",
			Description: "Administrative access excluding user management",
			Level:       4,
			Color:       "purple",
			IsSystem:    true,
			IsDefault:   false,
			IsActive:    true,
		},
		{
			Name:        "manager",
			DisplayName: "Manager",
			Description: "Management level access to products, orders, and customers",
			Level:       3,
			Color:       "blue",
			IsSystem:    false,
			IsDefault:   true,
			IsActive:    true,
		},
		{
			Name:        "warehouse_staff",
			DisplayName: "Warehouse Staff",
			Description: "Access to inventory and order processing",
			Level:       2,
			Color:       "green",
			IsSystem:    false,
			IsDefault:   false,
			IsActive:    true,
		},
		{
			Name:        "sales_rep",
			DisplayName: "Sales Representative",
			Description: "Customer and order management access",
			Level:       1,
			Color:       "orange",
			IsSystem:    false,
			IsDefault:   false,
			IsActive:    true,
		},
	}
}
