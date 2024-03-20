package domain

import (
	"time"

	"gorm.io/gorm"
)

type Narahubung struct {
	ID               uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	DetailKegiatanID uint           `gorm:"not null" json:"detail_kegiatan_id"`
	Judul            string         `gorm:"not null" json:"judul"`
	NamaNarahubung   string         `gorm:"not null" json:"nama_narahubung"`
	NoTelepon        string         `gorm:"not null" json:"no_telepon"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type NarahubungRepository interface {
	CreateNarahubung(req *Narahubung) (*Narahubung, error)
}
