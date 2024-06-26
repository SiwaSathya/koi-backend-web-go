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
	Email           string         `gorm:"not null" json:"email"`
	TipePembayaran  string         `gorm:"not null" json:"tipe_pembayaran"`
	NoTelepon       string         `gorm:"not null" json:"no_telepon"`
	Institusi       string         `gorm:"not null" json:"institusi"`
	BuktiPembayaran string         `gorm:"not null" json:"bukti_pembayaran"`
	Status          string         `gorm:"default:pending" json:"status"`
	Event           *Event         `json:"event"`
	Mahasiswa       *Mahasiswa     `json:"mahasiswa"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PembayaranRepository interface {
	CreatePembayaran(req *Pembayaran) (*Pembayaran, error)
	GetEvents() ([]Pembayaran, error)
	GetEventByMahasiswaID(id uint) ([]Pembayaran, error)
	UpdatePembayaran(req *Pembayaran) (*Pembayaran, error)
	UpdateStatusPembayaran(req *Pembayaran) (*Pembayaran, error)
	DeletePembayaran(id uint) error
}

type PembayaranUseCase interface {
	CreatePembayaran(ctx context.Context, req *Pembayaran) (*Pembayaran, error)
	GetEvents(ctx context.Context) ([]Pembayaran, error)
	GetEventByMahasiswaID(ctx context.Context, id uint) ([]Pembayaran, error)
	UpdatePembayaran(ctx context.Context, req *Pembayaran) (*Pembayaran, error)
	UpdateStatusPembayaran(ctx context.Context, req *Pembayaran) (*Pembayaran, error)
	DeletePembayaran(ctx context.Context, id uint) error
}
