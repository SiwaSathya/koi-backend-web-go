package domain

import (
	"time"

	"gorm.io/gorm"
)

type DetailKegiatan struct {
	ID               uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdEvent          uint           `gorm:"not null" json:"id_event"`
	MetodePembayaran string         `gorm:"not null" json:"metode_pembayaran"`
	HargaTiket       uint           `gorm:"not null" json:"harga_tiket"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
