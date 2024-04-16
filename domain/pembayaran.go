package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Pembayaran struct {
	ID              uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NamaPeserta     string         `gorm:"not null" json:"nama_peserta"`
	MahasiswaID     uint           `gorm:"mahasiswa_id" json:"mahasiswa_id"`
	EventID         uint           `gorm:"not null" json:"event_id"`
	NoTelepon       string         `gorm:"not null" json:"no_telepon"`
	Institusi       string         `gorm:"not null" json:"institusi"`
	BuktiPembayaran string         `gorm:"not null" json:"bukti_pembayaran"`
	Status          uint           `gorm:"default:0" json:"status"`
	Event           *Event         `json:"event"`
	Mahasiswa       *Mahasiswa     `json:"mahasiswa"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PembayaranRepository interface {
	CreatePembayaran(req *Pembayaran) (*Pembayaran, error)
	GetEventByMahasiswaID(id uint) ([]Pembayaran, error)
}

type PembayaranUseCase interface {
	CreatePembayaran(ctx context.Context, req *Pembayaran) (*Pembayaran, error)
	GetEventByMahasiswaID(ctx context.Context, id uint) ([]Pembayaran, error)
}
