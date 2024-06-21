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
	Category        string          `gorm:"not null" json:"category"`
	TanggalKegiatan string          `gorm:"not null" json:"tanggal_kegiatan"`
	TingkatKegiatan string          `gorm:"not null" json:"tingkat_kegiatan"`
	TypeImplement   string          `gorm:"not null" json:"type_implement"`
	Ormawa          *Ormawa         `json:"ormawa"`
	DetailKegiatan  *DetailKegiatan `json:"detail_kegiatan"`
	Pembayaran      *Pembayaran     `json:"pembayaran"`
	Absensi         *Absensi        `json:"absensi"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type CreateEvent struct {
	ID               uint               `json:"id"`
	OrmawaID         uint               `json:"ormawa_id"`
	NamaKegiatan     string             `json:"nama_kegiatan"`
	TanggalKegiatan  string             `json:"tanggal_kegiatan"`
	TingkatKegiatan  string             `json:"tingkat_kegiatan"`
	Category         string             `json:"category"`
	TypeImplement    string             `json:"type_implement"`
	HargaTiket       uint               `json:"harga_tiket"`
	DetailKegiatan   DetailKegiatan     `json:"detail_kegiatan"`
	MetodePembayaran []MetodePembayaran `json:"metode_pembayaran"`
	Narahubung       []Narahubung       `json:"narahubung"`
}

// // event
// id: event?.id || idParam,
// name_kegiatan: formState.name_event,
// harga_tiket: formState.harga_tiket,
// its_open: formState.status_submission,
// category: formState.category,
// tanggal_kegiatan: formState.tanggal_kegiatan,
// tingkat_kegiatan: formState.tingkat_kegiatan,
// type_implement: formState.type_implement,
// narahubung: contactPersons.map((contactPerson) => ({
// // id: 0,
// // detail_kegiatan_id: 0,
// judul: `Narahubung ${event?.nama_kegiatan} ${contactPerson.name}`,
// nama_narahubung: contactPerson.name,
// no_telepon: contactPerson.phoneNumber,
// })),
// metode_pembayaran: paymentMethods.map((paymentMethod) => ({
// // id: 0,
// // detail_kegiatan_id: 0,
// metode_pembayaran: paymentMethod.paymentMethod,
// judul: `Pembayaran Tiket ${event?.nama_kegiatan}`,
// nama_bank: paymentMethod.bankName,
// no_rekening: paymentMethod.accountNumber.toString(),
// pemilik: paymentMethod.ownerName,
// })),
// // detail event
// waktu_pelaksanaan: formState.waktu_pelaksanaan,
// lokasi: formState.lokasi,
// deskripsi: formState.deskripsi,
// gambar_kegiatan: formState.gambar_kegiatan,
// file_pengajuan: formState.file_pengajuan,
// // absensi
// its_close: formState.status_absensi,
// type UpdateEvent struct {
// 	ID               uint               `json:"id"`
// 	NameKegiatan     string             `json:"name_kegiatan"`
// 	HargaTiket       uint               `json:"harga_tiket"`
// 	ItsOpen          uint               `json:"its_open"`
// 	Category         string             `json:"category"`
// 	TanggalKegiatan  string             `json:"tanggal_kegiatan"`
// 	TingkatKegiatan  string             `json:"tingkat_kegiatan"`
// 	TypeImplement    string             `json:"type_implement"`
// 	Narahubung       []Narahubung       `json:"narahubung"`
// 	MetodePembayaran []MetodePembayaran `json:"metode_pembayaran"`
// 	// detail event
// 	WaktuPelaksanaan string `json:"waktu_pelaksanaan"`
// 	Lokasi           string `json:"lokasi"`
// 	Deskripsi        string `json:"deskripsi"`
// 	GambarKegiatan   string `json:"gambar_kegiatan"`
// 	FilePengajuan    string `json:"file_pengajuan"`
// 	// absensi
// 	ItsClose uint `json:"its_close"`
// }

type ResponeListEventOrmawa struct {
	Ormawa Ormawa  `json:"ormawa"`
	Event  []Event `json:"event"`
}

type EventRepository interface {
	CreateEvent(req *Event) (*Event, error)
	GetAllEvents() ([]Event, error)
	GetEventByOrmawaID(id uint) ([]Event, error)
	UpdateEvent(req *Event) error
	GetEventByID(id uint) (*Event, error)
	GetEventByIDAndOrmawaID(idOrmawa uint, idEvent uint) (*Event, error)
	DeleteEvent(id uint) error
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, req *CreateEvent) (*Event, error)
	GetAllEvents(ctx context.Context) ([]Event, error)
	GetEventByOrmawaID(ctx context.Context, id uint) (*ResponeListEventOrmawa, error)
	UpdateEvent(ctx context.Context, req *CreateEvent) error
	GetEventByID(ctx context.Context, id uint) (*Event, error)
	GetEventByIDAndOrmawaID(ctx context.Context, idUser uint, idEvent uint) (*Event, error)
	DeleteEvent(ctx context.Context, id uint) error
}
