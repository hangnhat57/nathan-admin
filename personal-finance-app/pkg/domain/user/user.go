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
	Email     string    `gorm:"type:varchar(255);unique"`
	Password  string    `gorm:"not null"`
	Role      Role      `gorm:"type:enum('admin', 'vip', 'normal');default:'normal'"`
	IsActive  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
