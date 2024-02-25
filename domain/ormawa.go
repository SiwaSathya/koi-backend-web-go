package domain

import (
	"time"

	"gorm.io/gorm"
)

type Ormawa struct {
	ID         uint           `gorm:"not null" json:"id"`
	NamaOrmawa string         `gorm:"not null" json:"nama_ormawa"`
	Status     int            `gorm:"not null" json:"status"`
	Password   string         `gorm:"not null" json:"password"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
