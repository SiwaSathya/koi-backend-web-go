package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	Nim           uint           `gorm:"primarykey;AUTO_INCREMENT" json:"nim"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	NoTelepon     string         `gorm:"not null" json:"no_telepon"`
	Email         string         `gorm:"not null" json:"email"`
	TanggalLahir  string         `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin  uint           `gorm:"not null" json:"jenis_kelamin"`
	TempatLahir   string         `gorm:"not null" json:"tempat_lahir"`
	AlamatTinggal string         `gorm:"not null" json:"alamat_tinggal"`
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
