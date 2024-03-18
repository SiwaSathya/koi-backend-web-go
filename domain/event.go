package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID              uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdOrmawa        uint           `gorm:"not null" json:"id_ormawa"`
	NamaKegiatan    string         `gorm:"not null" json:"nama_kegiatan"`
	TanggalKegiatan string         `gorm:"not null" json:"tanggal_kegiatan"`
	TingkatKegiatan string         `gorm:"not null" json:"tingkat_kegiatan"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
