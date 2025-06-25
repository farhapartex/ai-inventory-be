package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       string     `json:"employee_id" gorm:"uniqueIndex;not null;size:20"`
	FirstName        string     `json:"first_name" gorm:"size:150;not null" binding:"required"`
	LastName         string     `json:"last_name" gorm:"size:150;not null" binding:"required"`
	Email            string     `json:"email" gorm:"size:150;not null;unique" binding:"required,email"`
	Password         string     `json:"password" gorm:"size:255;not null"`
	Phone            string     `json:"phone" gorm:"size:20"`
	SecondaryPhone   string     `json:"secondary_phone" gorm:"size:20"`
	Address          string     `json:"address" gorm:"size:500"`
	City             string     `json:"city" gorm:"size:100"`
	State            string     `json:"state" gorm:"size:100"`
	ZipCode          string     `json:"zip_code" gorm:"size:20"`
	Country          string     `json:"country" gorm:"size:100;default:'United States'"`
	DateOfBirth      *time.Time `json:"date_of_birth"`
	Gender           string     `json:"gender" gorm:"size:20;"`
	DepartmentID     *uint      `json:"department_id" gorm:"index"`
	DepartmentRoleID *uint      `json:"department_role_id" gorm:"index"`
	ManagerID        *uint      `json:"manager_id" gorm:"index"`
	RoleID           *uint      `json:"role_id" gorm:"index"`
	JobTitle         string     `json:"job_title" gorm:"size:150"`
	WorkLocation     string     `json:"work_location" gorm:"size:150;default:'Main Office'"`
	ContractType     string     `json:"contract_type" gorm:"size:50;default:'Full-time';check:contract_type IN ('Full-time', 'Part-time', 'Contract', 'Intern', 'Consultant')"`
	EmploymentStatus string     `json:"employment_status" gorm:"size:50;default:'Active';check:employment_status IN ('Active', 'Inactive', 'Terminated', 'On Leave', 'Probation')"`
	HireDate         *time.Time `json:"hire_date"`
	TerminationDate  *time.Time `json:"termination_date"`
	ProbationEndDate *time.Time `json:"probation_end_date"`
	Salary           *float64   `json:"salary,omitempty" gorm:"type:decimal(12,2)"`
	Currency         string     `json:"currency" gorm:"size:3;default:'USD'"`
	PayGrade         string     `json:"pay_grade" gorm:"size:20"`
	IsSuperuser      bool       `gorm:"default:false" json:"is_superuser"`
	Status           string     `gorm:"size:20;default:active;" json:"status"`
	EmailVerified    bool       `gorm:"default:false" json:"email_verified"`
	VerifiedAt       *time.Time `json:"verified_at"`

	TwoFactorEnabled       bool       `json:"two_factor_enabled" gorm:"default:false"`
	TwoFactorSecret        string     `json:"two_factor_secret,omitempty" gorm:"size:255"`
	BackupCodes            string     `json:"backup_codes,omitempty" gorm:"type:text"`
	AccountLocked          bool       `json:"account_locked" gorm:"default:false"`
	LockedAt               *time.Time `json:"locked_at"`
	FailedLoginAttempts    int        `json:"failed_login_attempts" gorm:"default:0"`
	LastFailedLoginAt      *time.Time `json:"last_failed_login_at"`
	LoginCount             int        `json:"login_count" gorm:"default:0"`
	LastLoginAt            *time.Time `json:"last_login_at"`
	LastLoginIP            string     `json:"last_login_ip" gorm:"size:45"`
	LastPasswordChangedAt  *time.Time `json:"password_changed_at"`
	PasswordExpiresAt      *time.Time `json:"password_expires_at"`
	MustChangePassword     bool       `json:"must_change_password" gorm:"default:false"`
	TokenVersion           int        `json:"token_version" gorm:"default:1"` // For token invalidation
	LastTokenIssuedAt      *time.Time `json:"last_token_issued_at"`
	PreferredLanguage      string     `json:"preferred_language" gorm:"size:10;default:'en'"`
	PreferredCommunication string     `json:"preferred_communication" gorm:"size:20;default:'Email';check:preferred_communication IN ('Email', 'Phone', 'SMS')"`
	MarketingOptIn         bool       `json:"marketing_opt_in" gorm:"default:true"`
	NewsletterSubscription bool       `json:"newsletter_subscription" gorm:"default:false"`

	Notes         string `json:"notes" gorm:"type:text"`
	ReferredBy    *uint  `json:"referred_by" gorm:"index"`
	Source        string `json:"source" gorm:"size:100"`
	LoyaltyPoints int    `json:"loyalty_points" gorm:"default:0"`

	LinkedInProfile string `json:"linkedin_profile" gorm:"size:255"`
	TwitterHandle   string `json:"twitter_handle" gorm:"size:100"`
	Website         string `json:"website" gorm:"size:255"`

	EmergencyContactName     string `json:"emergency_contact_name" gorm:"size:200"`
	EmergencyContactPhone    string `json:"emergency_contact_phone" gorm:"size:20"`
	EmergencyContactRelation string `json:"emergency_contact_relation" gorm:"size:50"`

	CreatedBy       *uint            `json:"created_by" gorm:"index"`
	JoinedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `json:"deleted_at,omitempty" gorm:"index"`
	Role            *Role            `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	Department      *Department      `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	DepartmentRole  *DepartmentRole  `json:"department_role,omitempty" gorm:"foreignKey:DepartmentRoleID"`
	Manager         *User            `json:"manager,omitempty" gorm:"foreignKey:ManagerID"`
	DirectReports   []User           `json:"direct_reports,omitempty" gorm:"foreignKey:ManagerID"`
	Creator         *User            `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	ReferredByUser  *User            `json:"referred_by_user,omitempty" gorm:"foreignKey:ReferredBy"`
	UserHistory     []UserHistory    `json:"user_history,omitempty" gorm:"foreignKey:UserID"`
	UserPermissions []UserPermission `json:"user_permissions,omitempty" gorm:"foreignKey:UserID"`
	TokenBlacklist  []TokenBlacklist `json:"token_blacklist,omitempty" gorm:"foreignKey:UserID"`
}

type TokenBlacklist struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	TokenJTI  string    `json:"token_jti" gorm:"uniqueIndex;not null;size:255"` // JWT ID
	TokenHash string    `json:"token_hash" gorm:"index;size:64"`                // SHA256 hash of token
	Reason    string    `json:"reason" gorm:"size:100"`                         // logout, security, admin_revoke
	ExpiresAt time.Time `json:"expires_at" gorm:"index"`                        // When the original token would expire
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type UserHistory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Action    string    `json:"action" gorm:"not null;size:50"`
	Details   string    `json:"details" gorm:"type:text"`
	IPAddress string    `json:"ip_address" gorm:"size:45"`
	UserAgent string    `json:"user_agent" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type UserPermission struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	UserID        uint       `json:"user_id" gorm:"not null;index"`
	PermissionID  uint       `json:"permission_id" gorm:"not null;index"`
	IsGranted     bool       `json:"is_granted" gorm:"default:true"`
	GrantedBy     uint       `json:"granted_by" gorm:"not null;index"`
	GrantedAt     time.Time  `json:"granted_at" gorm:"autoCreateTime"`
	ExpiresAt     *time.Time `json:"expires_at"`
	Reason        string     `json:"reason" gorm:"size:300"`
	User          User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Permission    Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
	GrantedByUser User       `json:"granted_by_user,omitempty" gorm:"foreignKey:GrantedBy"`

	_ struct{} `gorm:"uniqueIndex:idx_user_permission,priority:1" sql:"user_id"`
	_ struct{} `gorm:"uniqueIndex:idx_user_permission,priority:2" sql:"permission_id"`
}

type UserProfile struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"uniqueIndex;not null"`
	Avatar         string    `json:"avatar" gorm:"size:500"`
	Bio            string    `json:"bio" gorm:"type:text"`
	Skills         string    `json:"skills" gorm:"type:text"`         // JSON array of skills
	Certifications string    `json:"certifications" gorm:"type:text"` // JSON array
	Education      string    `json:"education" gorm:"type:text"`      // JSON array
	Experience     string    `json:"experience" gorm:"type:text"`     // JSON array
	Interests      string    `json:"interests" gorm:"type:text"`      // JSON array
	Timezone       string    `json:"timezone" gorm:"size:50;default:'UTC'"`
	WorkingHours   string    `json:"working_hours" gorm:"type:text"` // JSON object
	PublicProfile  bool      `json:"public_profile" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// BeforeCreate hook for User
func (u *User) BeforeCreate(tx *gorm.DB) error {

	if u.EmployeeID == "" {
		u.EmployeeID = generateEmployeeID(tx)
	}

	if len(u.Password) < 60 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	return nil
}

// func (u *User) AfterCreate(tx *gorm.DB) error {
// 	return tx.Model(&Department{}).Where("id = ?", u.DepartmentID).
// 		Update("employee_count", gorm.Expr("employee_count + 1")).Error
// }

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" && len(u.Password) < 60 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
		now := time.Now()
		u.LastPasswordChangedAt = &now
	}

	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	return tx.Model(&Department{}).Where("id = ?", u.DepartmentID).
		Update("employee_count", gorm.Expr("employee_count - 1")).Error
}

// User methods
func (u *User) IsPasswordExpired() bool {
	if u.PasswordExpiresAt == nil {
		return false
	}
	return time.Now().After(*u.PasswordExpiresAt)
}

// ShouldChangePassword checks if user should be forced to change password
func (u *User) ShouldChangePassword() bool {
	return u.MustChangePassword || u.IsPasswordExpired()
}

// ResetFailedLoginAttempts resets failed login attempts
func (u *User) ResetFailedLoginAttempts(tx *gorm.DB) error {
	u.FailedLoginAttempts = 0
	u.LastFailedLoginAt = nil
	return tx.Save(u).Error
}

// RecordLogin records a successful login
func (u *User) RecordLogin(tx *gorm.DB, ipAddress string) error {
	now := time.Now()
	u.LastLoginAt = &now
	u.LastLoginIP = ipAddress
	u.LoginCount++

	// Reset failed attempts on successful login
	u.FailedLoginAttempts = 0
	u.LastFailedLoginAt = nil

	return tx.Save(u).Error
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) HasPermission(tx *gorm.DB, permissionName string) bool {
	if u.Role.HasPermission(tx, permissionName) {
		return true
	}
	var count int64
	tx.Table("user_permissions").
		Joins("JOIN permissions ON permissions.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ? AND permissions.name = ? AND user_permissions.is_granted = ? AND permissions.is_active = ?",
			u.ID, permissionName, true, true).
		Where("user_permissions.expires_at IS NULL OR user_permissions.expires_at > ?", time.Now()).
		Count(&count)

	return count > 0
}

func (u *User) IncrementTokenVersion(tx *gorm.DB) error {
	u.TokenVersion++
	return tx.Save(u).Error
}

// RecordTokenIssue records when a token was issued
func (u *User) RecordTokenIssue(tx *gorm.DB, ipAddress string) error {
	now := time.Now()
	u.LastTokenIssuedAt = &now
	u.LastLoginAt = &now
	u.LastLoginIP = ipAddress
	u.LoginCount++

	// Reset failed attempts on successful login
	u.FailedLoginAttempts = 0
	u.LastFailedLoginAt = nil

	return tx.Save(u).Error
}

// BlacklistToken adds a token to the blacklist (for logout/revocation)
func (u *User) BlacklistToken(tx *gorm.DB, tokenJTI, tokenHash string, expiresAt time.Time, reason string) error {
	blacklistedToken := TokenBlacklist{
		UserID:    u.ID,
		TokenJTI:  tokenJTI,
		TokenHash: tokenHash,
		Reason:    reason,
		ExpiresAt: expiresAt,
	}
	return tx.Create(&blacklistedToken).Error
}

// IsTokenBlacklisted checks if a token is blacklisted
func (u *User) IsTokenBlacklisted(tx *gorm.DB, tokenJTI string) bool {
	var count int64
	tx.Model(&TokenBlacklist{}).
		Where("user_id = ? AND token_jti = ? AND expires_at > ?", u.ID, tokenJTI, time.Now()).
		Count(&count)
	return count > 0
}

// LogoutAllSessions invalidates all user tokens by incrementing version
func (u *User) LogoutAllSessions(tx *gorm.DB, reason string) error {
	// Increment token version to invalidate all tokens
	if err := u.IncrementTokenVersion(tx); err != nil {
		return err
	}

	// Log the action
	history := UserHistory{
		UserID:  u.ID,
		Action:  "LOGOUT_ALL",
		Details: fmt.Sprintf(`{"reason": "%s"}`, reason),
	}
	return tx.Create(&history).Error
}

// Other methods remain the same...
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	now := time.Now()
	u.LastPasswordChangedAt = &now
	return nil
}

func (u *User) IncrementFailedLoginAttempts(tx *gorm.DB) error {
	u.FailedLoginAttempts++
	now := time.Now()
	u.LastFailedLoginAt = &now

	// Lock account after 5 failed attempts
	if u.FailedLoginAttempts >= 5 {
		u.AccountLocked = true
		u.LockedAt = &now
	}

	return tx.Save(u).Error
}

func (u *User) UnlockAccount(tx *gorm.DB) error {
	u.AccountLocked = false
	u.LockedAt = nil
	u.FailedLoginAttempts = 0
	u.LastFailedLoginAt = nil
	return tx.Save(u).Error
}

func (u *User) IsActive() bool {
	return u.Status == "active" && !u.AccountLocked
}

func (u *User) CanLogin() bool {
	return u.IsActive() && u.EmailVerified
}

// JWT Claims structure (for reference)
type JWTClaims struct {
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
	RoleID       uint   `json:"role_id"`
	DepartmentID uint   `json:"department_id"`
	TokenVersion int    `json:"token_version"`
	TokenType    string `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

func generateUserID(tx *gorm.DB) string {
	for {
		userID := fmt.Sprintf("USR-%04d", time.Now().Unix()%10000)
		var count int64
		tx.Model(&User{}).Where("user_id = ?", userID).Count(&count)
		if count == 0 {
			return userID
		}
	}
}

func generateEmployeeID(tx *gorm.DB) string {
	for {
		empID := fmt.Sprintf("EMP-%04d", time.Now().Unix()%10000)
		var count int64
		tx.Model(&User{}).Where("employee_id = ?", empID).Count(&count)
		if count == 0 {
			return empID
		}
	}
}
