package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	Nim       uint           `gorm:"not null" json:"nim"`
	Username  string         `gorm:"not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
