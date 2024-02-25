package domain

import (
	"time"

	"gorm.io/gorm"
)

type DetailKegiatan struct {
	ID               uint           `gorm:"not null" json:"id"`
	IdEvent          string         `gorm:"not null" json:"id_event"`
	IdOrmawa         string         `gorm:"not null" json:"id_ormawa"`
	NamaKegiatan     string         `gorm:"not null" json:"nama_kegiatan"`
	TanggalKegiatan  string         `gorm:"not null" json:"tanggal_kegiatan"`
	MetodePembayaran string         `gorm:"not null" json:"metode_pembayaran"`
	HargaTiket       uint           `gorm:"not null" json:"harga_tiket"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
