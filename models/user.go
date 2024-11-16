package models

import (
	"time"
)

type UserRole string
type AuthMethod string

const (
	RoleStudent  UserRole   = "STUDENT"
	RoleAdmin    UserRole   = "ADMIN"
	MethodCreds  AuthMethod = "CREDENTIALS"
	MethodGoogle AuthMethod = "GOOGLE"
	MethodGithub AuthMethod = "GITHUB"
)

type User struct {
	ID            string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email         string     `gorm:"unique;not null" json:"email"`
	Password      *string    `gorm:"type:text" json:"password,omitempty"`
	Username      string     `gorm:"unique;not null" json:"username"`
	DisplayName   string     `gorm:"column:display_name;not null" json:"display_name"`
	Picture       *string    `json:"picture,omitempty"`
	Points        int        `gorm:"default:0" json:"points"`
	Method        AuthMethod `gorm:"type:text" json:"method"`
	Role          UserRole   `gorm:"type:text;default:STUDENT" json:"role"`
	IsVerified    bool       `gorm:"column:is_verified;default:false" json:"is_verified"`
	IsTotpEnabled bool       `gorm:"column:is_totp_enabled;default:false" json:"is_totp_enabled"`
	TotpSecret    string     `gorm:"default:'totp_secret'" json:"totp_secret"`
	CreatedAt     time.Time  `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
