package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	Nim       uint           `gorm:"primarykey;AUTO_INCREMENT" json:"nim"`
	Username  string         `gorm:"not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `json:"user"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type MahasiswaRepository interface {
	CreateMahasiswa(req *Mahasiswa) (*Mahasiswa, error)
}
