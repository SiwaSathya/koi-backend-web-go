package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	Nim           uint           `gorm:"primarykey;AUTO_INCREMENT" json:"nim"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	NoTelepon     *string        `gorm:"null" json:"no_telepon"`
	NamaMahasiswa *string        `gorm:"null" json:"nama_mahasiswa"`
	Email         *string        `gorm:"null" json:"email"`
	TanggalLahir  *string        `gorm:"null" json:"tanggal_lahir"`
	JenisKelamin  *uint          `gorm:"null" json:"jenis_kelamin" default:"0"`
	TempatLahir   *string        `gorm:"null" json:"tempat_lahir"`
	AlamatTinggal *string        `gorm:"null" json:"alamat_tinggal"`
	Photo         *string        `gorm:"null" json:"photo"`
	User          *User          `json:"user"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type MahasiswaRepository interface {
	CreateMahasiswa(req *Mahasiswa) (*Mahasiswa, error)
	GetMahasiswaByUserID(userID uint) (*Mahasiswa, error)
	UpdateMahasiswa(req *Mahasiswa) error
}
