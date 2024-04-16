package domain

import (
	"time"

	"gorm.io/gorm"
)

type MetodePembayaran struct {
	ID               uint            `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	DetailKegiatanID uint            `gorm:"not null" json:"detail_kegiatan_id"`
	Judul            string          `gorm:"not null" json:"judul"`
	NamaBank         string          `gorm:"not null" json:"nama_bank"`
	NoRekening       string          `gorm:"not null" json:"no_rekening"`
	PemilikRekening  string          `gorm:"not null" json:"pemilik"`
	DetailKegiatan   *DetailKegiatan `gorm:"not null" json:"detail_kegiatan"`
	CreatedAt        *time.Time      `json:"created_at"`
	UpdatedAt        *time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type MetodePembayaranRepository interface {
	CreateMetodePembayaran(req *MetodePembayaran) (*MetodePembayaran, error)
}
