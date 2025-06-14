package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string         `json:"name" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	Code            string         `json:"code" gorm:"uniqueIndex;not null;size:20" binding:"required"`
	Description     string         `json:"description" gorm:"size:500"`
	ParentID        *uint          `json:"parent_id" gorm:"index"`
	ManagerID       *uint          `json:"manager_id" gorm:"index"`
	BudgetCode      string         `json:"budget_code" gorm:"size:50;index"`
	CostCenter      string         `json:"cost_center" gorm:"size:50;index"`
	Location        string         `json:"location" gorm:"size:200"`
	PhoneNumber     string         `json:"phone_number" gorm:"size:20"`
	Email           string         `json:"email" gorm:"size:100"`
	IsActive        bool           `json:"is_active" gorm:"default:true"`
	EmployeeCount   int            `json:"employee_count" gorm:"default:0"`
	MaxEmployees    *int           `json:"max_employees"`
	EstablishedDate *time.Time     `json:"established_date"`
	CreatedBy       uint           `json:"created_by" gorm:"index"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Parent          *Department      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children        []Department     `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Manager         *User            `json:"manager,omitempty" gorm:"foreignKey:ManagerID"`
	Users           []User           `json:"users,omitempty" gorm:"foreignKey:DepartmentID"`
	Creator         *User            `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	DepartmentRoles []DepartmentRole `json:"department_roles,omitempty" gorm:"foreignKey:DepartmentID"`
}

type DepartmentRole struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DepartmentID uint      `json:"department_id" gorm:"not null;index"`
	Name         string    `json:"name" gorm:"not null;size:100" binding:"required"`
	Description  string    `json:"description" gorm:"size:300"`
	Level        int       `json:"level" gorm:"default:1;check:level >= 1 AND level <= 10"`
	IsManagerial bool      `json:"is_managerial" gorm:"default:false"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Department Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	Users      []User     `json:"users,omitempty" gorm:"foreignKey:DepartmentRoleID"`

	_ struct{} `gorm:"uniqueIndex:idx_dept_role_name,priority:1" sql:"department_id"`
	_ struct{} `gorm:"uniqueIndex:idx_dept_role_name,priority:2" sql:"name"`
}

type DepartmentHierarchy struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	AncestorID   uint `json:"ancestor_id" gorm:"not null;index"`
	DescendantID uint `json:"descendant_id" gorm:"not null;index"`
	Depth        int  `json:"depth" gorm:"not null;default:0"`

	Ancestor   Department `json:"ancestor,omitempty" gorm:"foreignKey:AncestorID"`
	Descendant Department `json:"descendant,omitempty" gorm:"foreignKey:DescendantID"`

	_ struct{} `gorm:"uniqueIndex:idx_dept_hierarchy,priority:1" sql:"ancestor_id"`
	_ struct{} `gorm:"uniqueIndex:idx_dept_hierarchy,priority:2" sql:"descendant_id"`
}

type DepartmentBudget struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	DepartmentID      uint       `json:"department_id" gorm:"not null;index"`
	FiscalYear        int        `json:"fiscal_year" gorm:"not null;index"`
	TotalBudget       float64    `json:"total_budget" gorm:"type:decimal(15,2);default:0"`
	OperationalBudget float64    `json:"operational_budget" gorm:"type:decimal(15,2);default:0"`
	CapitalBudget     float64    `json:"capital_budget" gorm:"type:decimal(15,2);default:0"`
	SpentAmount       float64    `json:"spent_amount" gorm:"type:decimal(15,2);default:0"`
	Currency          string     `json:"currency" gorm:"size:3;default:'USD'"`
	IsApproved        bool       `json:"is_approved" gorm:"default:false"`
	ApprovedBy        *uint      `json:"approved_by" gorm:"index"`
	ApprovedAt        *time.Time `json:"approved_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	Department     Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	ApprovedByUser *User      `json:"approved_by_user,omitempty" gorm:"foreignKey:ApprovedBy"`

	_ struct{} `gorm:"uniqueIndex:idx_dept_budget_year,priority:1" sql:"department_id"`
	_ struct{} `gorm:"uniqueIndex:idx_dept_budget_year,priority:2" sql:"fiscal_year"`
}

type DepartmentPermission struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DepartmentID uint      `json:"department_id" gorm:"not null;index"`
	PermissionID uint      `json:"permission_id" gorm:"not null;index"`
	IsGranted    bool      `json:"is_granted" gorm:"default:true"`
	GrantedBy    uint      `json:"granted_by" gorm:"not null;index"`
	GrantedAt    time.Time `json:"granted_at" gorm:"autoCreateTime"`
	Reason       string    `json:"reason" gorm:"size:300"`

	Department    Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	Permission    Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
	GrantedByUser User       `json:"granted_by_user,omitempty" gorm:"foreignKey:GrantedBy"`

	_ struct{} `gorm:"uniqueIndex:idx_dept_permission,priority:1" sql:"department_id"`
	_ struct{} `gorm:"uniqueIndex:idx_dept_permission,priority:2" sql:"permission_id"`
}

type DepartmentHistory struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DepartmentID uint      `json:"department_id" gorm:"not null;index"`
	Action       string    `json:"action" gorm:"not null;size:50"`
	FieldName    string    `json:"field_name" gorm:"size:50"`
	OldValue     string    `json:"old_value" gorm:"type:text"`
	NewValue     string    `json:"new_value" gorm:"type:text"`
	ChangedBy    uint      `json:"changed_by" gorm:"not null;index"`
	Reason       string    `json:"reason" gorm:"size:500"`
	IPAddress    string    `json:"ip_address" gorm:"size:45"`
	UserAgent    string    `json:"user_agent" gorm:"size:500"`
	CreatedAt    time.Time `json:"created_at"`

	Department    Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	ChangedByUser User       `json:"changed_by_user,omitempty" gorm:"foreignKey:ChangedBy"`
}

// Model methods and hooks

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	// Generate unique code if not provided
	if d.Code == "" {
		d.Code = generateDepartmentCode(tx, d.Name)
	}
	return nil
}

// AfterCreate hook for Department
func (d *Department) AfterCreate(tx *gorm.DB) error {
	return d.createHierarchyEntries(tx)
}

// AfterUpdate hook for Department
func (d *Department) AfterUpdate(tx *gorm.DB) error {
	return d.UpdateEmployeeCount(tx)
}

// BeforeDelete hook for Department
func (d *Department) BeforeDelete(tx *gorm.DB) error {
	var userCount int64
	tx.Model(&User{}).Where("department_id = ? AND is_active = ?", d.ID, true).Count(&userCount)
	if userCount > 0 {
		return gorm.ErrRecordNotFound
	}

	var childCount int64
	tx.Model(&Department{}).Where("parent_id = ?", d.ID).Count(&childCount)
	if childCount > 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (d *Department) UpdateEmployeeCount(tx *gorm.DB) error {
	var count int64
	tx.Model(&User{}).Where("department_id = ? AND is_active = ?", d.ID, true).Count(&count)
	return tx.Model(d).Update("employee_count", count).Error
}

func (d *Department) GetAllAncestors(tx *gorm.DB) ([]Department, error) {
	var ancestors []Department
	err := tx.Table("departments").
		Joins("JOIN department_hierarchies ON departments.id = department_hierarchies.ancestor_id").
		Where("department_hierarchies.descendant_id = ? AND department_hierarchies.depth > 0", d.ID).
		Order("department_hierarchies.depth DESC").
		Find(&ancestors).Error
	return ancestors, err
}

// GetAllDescendants returns all descendant departments
func (d *Department) GetAllDescendants(tx *gorm.DB) ([]Department, error) {
	var descendants []Department
	err := tx.Table("departments").
		Joins("JOIN department_hierarchies ON departments.id = department_hierarchies.descendant_id").
		Where("department_hierarchies.ancestor_id = ? AND department_hierarchies.depth > 0", d.ID).
		Order("department_hierarchies.depth ASC").
		Find(&descendants).Error
	return descendants, err
}

// GetDirectChildren returns immediate child departments
func (d *Department) GetDirectChildren(tx *gorm.DB) ([]Department, error) {
	var children []Department
	err := tx.Where("parent_id = ?", d.ID).Find(&children).Error
	return children, err
}

// IsAncestorOf checks if this department is an ancestor of another department
func (d *Department) IsAncestorOf(tx *gorm.DB, departmentID uint) bool {
	var count int64
	tx.Table("department_hierarchies").
		Where("ancestor_id = ? AND descendant_id = ? AND depth > 0", d.ID, departmentID).
		Count(&count)
	return count > 0
}

// createHierarchyEntries creates hierarchy entries for department tree structure
func (d *Department) createHierarchyEntries(tx *gorm.DB) error {
	// Self reference with depth 0
	hierarchy := DepartmentHierarchy{
		AncestorID:   d.ID,
		DescendantID: d.ID,
		Depth:        0,
	}
	if err := tx.Create(&hierarchy).Error; err != nil {
		return err
	}

	// If has parent, create entries for all ancestors
	if d.ParentID != nil {
		var parentHierarchies []DepartmentHierarchy
		err := tx.Where("descendant_id = ?", *d.ParentID).Find(&parentHierarchies).Error
		if err != nil {
			return err
		}

		for _, ph := range parentHierarchies {
			newHierarchy := DepartmentHierarchy{
				AncestorID:   ph.AncestorID,
				DescendantID: d.ID,
				Depth:        ph.Depth + 1,
			}
			if err := tx.Create(&newHierarchy).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// Helper functions

// generateDepartmentCode generates a unique department code
func generateDepartmentCode(tx *gorm.DB, name string) string {
	// Simple implementation - you can make this more sophisticated
	baseCode := ""
	words := strings.Fields(strings.ToUpper(name))
	for i, word := range words {
		if i < 3 { // Take first 3 words
			if len(word) > 0 {
				baseCode += string(word[0])
			}
		}
	}

	// Ensure uniqueness
	code := baseCode
	counter := 1
	for {
		var count int64
		tx.Model(&Department{}).Where("code = ?", code).Count(&count)
		if count == 0 {
			break
		}
		code = fmt.Sprintf("%s%d", baseCode, counter)
		counter++
	}

	return code
}

func GetDefaultDepartments() []Department {
	return []Department{
		{
			Name:        "Administration",
			Code:        "ADM",
			Description: "Administrative and executive functions",
			Location:    "Main Office",
			IsActive:    true,
		},
		{
			Name:        "Information Technology",
			Code:        "IT",
			Description: "Technology infrastructure and support",
			Location:    "Main Office",
			IsActive:    true,
		},
		{
			Name:        "Sales",
			Code:        "SAL",
			Description: "Sales and customer acquisition",
			Location:    "Main Office",
			IsActive:    true,
		},
		{
			Name:        "Warehouse Operations",
			Code:        "WH",
			Description: "Inventory management and logistics",
			Location:    "Warehouse",
			IsActive:    true,
		},
		{
			Name:        "Customer Service",
			Code:        "CS",
			Description: "Customer support and relations",
			Location:    "Main Office",
			IsActive:    true,
		},
		{
			Name:        "Finance",
			Code:        "FIN",
			Description: "Financial planning and accounting",
			Location:    "Main Office",
			IsActive:    true,
		},
	}
}

func GetDefaultDepartmentRoles() map[string][]DepartmentRole {
	return map[string][]DepartmentRole{
		"Administration": {
			{Name: "Executive", Description: "Executive level position", Level: 10, IsManagerial: true},
			{Name: "Manager", Description: "Department manager", Level: 8, IsManagerial: true},
			{Name: "Supervisor", Description: "Team supervisor", Level: 6, IsManagerial: true},
			{Name: "Administrator", Description: "Administrative staff", Level: 4, IsManagerial: false},
		},
		"Information Technology": {
			{Name: "IT Director", Description: "IT department head", Level: 9, IsManagerial: true},
			{Name: "Senior Developer", Description: "Senior development role", Level: 7, IsManagerial: false},
			{Name: "System Administrator", Description: "System administration", Level: 6, IsManagerial: false},
			{Name: "IT Support", Description: "Technical support staff", Level: 4, IsManagerial: false},
		},
		"Sales": {
			{Name: "Sales Director", Description: "Sales department head", Level: 9, IsManagerial: true},
			{Name: "Sales Manager", Description: "Sales team manager", Level: 7, IsManagerial: true},
			{Name: "Senior Sales Rep", Description: "Senior sales representative", Level: 6, IsManagerial: false},
			{Name: "Sales Representative", Description: "Sales representative", Level: 4, IsManagerial: false},
		},
		"Warehouse Operations": {
			{Name: "Warehouse Manager", Description: "Warehouse operations manager", Level: 8, IsManagerial: true},
			{Name: "Warehouse Supervisor", Description: "Warehouse team supervisor", Level: 6, IsManagerial: true},
			{Name: "Warehouse Staff", Description: "Warehouse operations staff", Level: 4, IsManagerial: false},
			{Name: "Inventory Clerk", Description: "Inventory management clerk", Level: 3, IsManagerial: false},
		},
		"Customer Service": {
			{Name: "CS Manager", Description: "Customer service manager", Level: 8, IsManagerial: true},
			{Name: "Senior CS Rep", Description: "Senior customer service rep", Level: 6, IsManagerial: false},
			{Name: "CS Representative", Description: "Customer service representative", Level: 4, IsManagerial: false},
		},
		"Finance": {
			{Name: "Finance Director", Description: "Finance department head", Level: 9, IsManagerial: true},
			{Name: "Senior Accountant", Description: "Senior accounting role", Level: 7, IsManagerial: false},
			{Name: "Accountant", Description: "Staff accountant", Level: 5, IsManagerial: false},
			{Name: "Finance Clerk", Description: "Finance support clerk", Level: 3, IsManagerial: false},
		},
	}
}
