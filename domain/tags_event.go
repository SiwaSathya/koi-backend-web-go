package domain

import (
	"time"

	"gorm.io/gorm"
)

type TagsEvent struct {
	ID              uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdEvent         uint           `gorm:"not null" json:"id_event"`
	IdDetalKegiatan uint           `gorm:"not null" json:"id_detail_kegiatan"`
	Tags            string         `gorm:"not null" json:"tags"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
