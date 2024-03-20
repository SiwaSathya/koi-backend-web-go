package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID              uint            `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	OrmawaID        uint            `gorm:"not null" json:"ormawa_id"`
	NamaKegiatan    string          `gorm:"not null" json:"nama_kegiatan"`
	HargaTiket      uint            `gorm:"not null" json:"harga_tiket"`
	ItsOpen         uint            `gorm:"default:0" json:"its_open"`
	TanggalKegiatan string          `gorm:"not null" json:"tanggal_kegiatan"`
	TingkatKegiatan string          `gorm:"not null" json:"tingkat_kegiatan"`
	Ormawa          *Ormawa         `gorm:"not null" json:"ormawa"`
	DetailKegiatan  *DetailKegiatan `gorm:"not null" json:"detail_kegiatan"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type CreateEvent struct {
	OrmawaID         uint               `json:"ormawa_id"`
	NamaKegiatan     string             `json:"nama_kegiatan"`
	TanggalKegiatan  string             `json:"tanggal_kegiatan"`
	TingkatKegiatan  string             `json:"tingkat_kegiatan"`
	HargaTiket       uint               `json:"harga_tiket"`
	DetailKegiatan   DetailKegiatan     `json:"detail_kegiatan"`
	MetodePembayaran []MetodePembayaran `json:"metode_pembayaran"`
	Narahubung       []Narahubung       `json:"narahubung"`
}

type ResponeListEventOrmawa struct {
	Ormawa Ormawa  `json:"ormawa"`
	Event  []Event `json:"event"`
}

type EventRepository interface {
	CreateEvent(req *Event) (*Event, error)
	GetAllEvents() ([]Event, error)
	GetEventByOrmawaID(id uint) ([]Event, error)
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *CreateEvent) (*Event, error)
	GetAllEvents(ctx context.Context) ([]Event, error)
	GetEventByOrmawaID(ctx context.Context, id uint) (*ResponeListEventOrmawa, error)
}
