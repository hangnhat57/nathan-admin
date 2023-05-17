package user

import (
	"time"
)

type Role string

const (
	Admin  Role = "admin"
	VIP    Role = "vip"
	Normal Role = "normal"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	Role      Role      `gorm:"type:enum('admin', 'vip', 'normal');default:'normal'"`
	IsActive  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SignUpInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin vip normal"`
}
