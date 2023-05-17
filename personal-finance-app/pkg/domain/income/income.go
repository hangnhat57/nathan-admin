package income

import (
	"time"
)

type Income struct {
	ID          uint      `gorm:"primaryKey"`
	UserId      uint      `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Description string    `gorm:"not null"`
	Source      string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
