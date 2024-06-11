package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	ID              uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	EventId         uint           `gorm:"not null" json:"event_id"`
	UserId          *uint          `gorm:"null" json:"user_id"`
	NamaMahasiswa   uint           `gorm:"not null" json:"name"`
	NoTelepon       string         `gorm:"not null" json:"no_telepon"`
	Institusi       uint           `gorm:"not null" json:"institusi"`
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
}

type AbsensiUseCase interface {
	CreateAbsensi(ctx context.Context, req *Absensi) (*Absensi, error)
}
