package domain

import (
	"time"

	"gorm.io/gorm"
)

type Ormawa struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NamaOrmawa  string         `gorm:"not null" json:"nama_ormawa"`
	Status      int            `gorm:"not null" json:"status"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Email       string         `gorm:"not null" json:"email"`
	Deskripsi   string         `gorm:"not null" json:"deskripsi"`
	JenisOrmawa string         `gorm:"not null" json:"jenis_ormawa"`
	Logo        string         `gorm:"not null" json:"logo"`
	Cover       string         `gorm:"not null" json:"cover"`
	User        *User          `json:"user"`
	Event       *Event         `json:"event"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type OrmawaRepository interface {
	CreateOrmawa(req *Ormawa) (*Ormawa, error)
	GetOrmawaByID(id uint) (*Ormawa, error)
	Updateormawa(req *Ormawa) error
}
