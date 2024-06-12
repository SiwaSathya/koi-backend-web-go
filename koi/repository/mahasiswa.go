package repository

import (
	"errors"
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreMahasiswaRepository struct {
	DB *gorm.DB
}

func NewPostgreMahasiswa(client *gorm.DB) domain.MahasiswaRepository {
	return &posgreMahasiswaRepository{
		DB: client,
	}
}

func (a *posgreMahasiswaRepository) CreateMahasiswa(req *domain.Mahasiswa) (*domain.Mahasiswa, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Mahasiswa{}, err
	}

	createdMahasiswa := &domain.Mahasiswa{}
	err = a.DB.
		Last(createdMahasiswa).
		Error

	if err != nil {
		return &domain.Mahasiswa{}, err
	}

	return createdMahasiswa, nil
}

func (a *posgreMahasiswaRepository) GetMahasiswaByUserID(userID uint) (*domain.Mahasiswa, error) {
	var res domain.Mahasiswa
	err := a.DB.
		Model(domain.Mahasiswa{}).
		Where("user_id = ?", userID).
		Find(&res).Error
	if err != nil {
		return &domain.Mahasiswa{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Mahasiswa{}, fmt.Errorf("record not found")
	}

	return &res, nil

}

func (a *posgreMahasiswaRepository) UpdateMahasiswa(req *domain.Mahasiswa) error {
	err := a.DB.
		Model(domain.Mahasiswa{}).
		Where("user_id = ?", req.UserID).
		Select("no_telepon", "email", "tanggal_lahir", "jenis_kelamin", "tempat_lahir", "alamat_tinggal", "nama_mahasiswa").
		Updates(map[string]interface{}{
			"nama_mahasiswa": req.NamaMahasiswa,
			"no_telepon":     req.NoTelepon,
			"email":          req.Email,
			"tanggal_lahir":  req.TanggalLahir,
			"jenis_kelamin":  req.JenisKelamin,
			"tempat_lahir":   req.TempatLahir,
			"alamat_tinggal": req.AlamatTinggal,
		}).
		Error

	if err != nil {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("record not found")
	}

	return nil
}
