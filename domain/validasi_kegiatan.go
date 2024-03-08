package domain

import (
	"time"

	"gorm.io/gorm"
)

type ValidasiKegiatan struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdMahasiswa uint           `gorm:"not null" json:"id_mhs"`
	Status      uint           `gorm:"not null" json:"status"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
