package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DetailKegiatan struct {
	ID               uint                `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	EventID          uint                `gorm:"not null" json:"event_id"`
	WaktuPelaksanaan string              `gorm:"not null" json:"waktu_pelaksanaan"`
	Lokasi           string              `gorm:"not null" json:"lokasi"`
	Status           string              `gorm:"default:'pending'" json:"status"`
	Deskripsi        string              `gorm:"not null" json:"deskripsi"`
	GambarKegiatan   string              `gorm:"not null" json:"gambar_kegiatan"`
	FilePengajuan    string              `gorm:"not null" json:"file_pengajuan"`
	Sertifikat       *string             `json:"sertifikat"`
	Event            *Event              `json:"event"`
	MetodePembayaran *[]MetodePembayaran `json:"metode_pembayaran"`
	Narahubung       *[]Narahubung       `json:"narahubung"`
	CreatedAt        *time.Time          `json:"created_at"`
	UpdatedAt        *time.Time          `json:"updated_at"`
	DeletedAt        gorm.DeletedAt      `gorm:"index" json:"deleted_at"`
}

type DetailKegiatanRepository interface {
	CreateDetailKegiatan(req *DetailKegiatan) (*DetailKegiatan, error)
	GetDetailKegiatanByID(id uint) (*DetailKegiatan, error)
	UpdateStatus(eventID uint, status string) error
	UpdateDetailKegiatan(req *DetailKegiatan) error
}

type DetailKegiatanUseCase interface {
	GetDetailKegiatanByID(ctx context.Context, id uint) (*DetailKegiatan, error)
	UpdateStatus(ctx context.Context, eventID uint, status string) error
}
