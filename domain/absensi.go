package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	ID              uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	EventId         uint           `gorm:"not null" json:"event_id"`
	UserId          uint           `gorm:"not null" json:"user_id"`
	NamaMahasiswa   string         `gorm:"not null" json:"name_mahasiswa"`
	NoTelepon       string         `gorm:"not null" json:"no_telepon"`
	Institusi       string         `gorm:"not null" json:"institusi"`
	Status          string         `gorm:"not null" json:"status" default:"pending"`
	BuktiPembayaran string         `gorm:"not null" json:"bukti_pembayaran"`
	ItsClose        uint           `gorm:"default:0" json:"its_close"`
	Category        string         `gorm:"not null" json:"category"`
	TanggalKegiatan string         `gorm:"not null" json:"tanggal_kegiatan"`
	TingkatKegiatan string         `gorm:"not null" json:"tingkat_kegiatan"`
	Event           *Event         `json:"event"`
	User            *User          `json:"user"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type AbsensiRepoository interface {
	CreateAbsensi(req *Absensi) (*Absensi, error)
	GetAllAbsensi() ([]Absensi, error)
	GetAbsensiByEventID(eventId uint) ([]Absensi, error)
	UpdateStatus(eventId uint, userId uint, status string) error
}

type AbsensiUseCase interface {
	CreateAbsensi(ctx context.Context, req *Absensi) (*Absensi, error)
	GetAllAbsensi(ctx context.Context) ([]Absensi, error)
	GetAbsensiByEventID(ctx context.Context, eventId uint) ([]Absensi, error)
	UpdateStatus(ctx context.Context, eventId uint, userId uint, status string) error
}
