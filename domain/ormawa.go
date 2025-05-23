package domain

import (
	"time"

	"gorm.io/gorm"
)

type Ormawa struct {
	ID         uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NamaOrmawa string         `gorm:"not null" json:"nama_ormawa"`
	Status     int            `gorm:"not null" json:"status"`
	Password   string         `gorm:"not null" json:"password"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	User       User           `json:"user"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type OrmawaRepository interface {
	CreateOrmawa(req *Ormawa) (*Ormawa, error)
}
