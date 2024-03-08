package repository

import (
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
