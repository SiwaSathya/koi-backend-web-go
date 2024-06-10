package repository

import (
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreAbsensiRepository struct {
	DB *gorm.DB
}

func NewPostgreAbsensi(client *gorm.DB) domain.AbsensiRepoository {
	return &posgreAbsensiRepository{
		DB: client,
	}
}

func (a *posgreAbsensiRepository) CreateAbsensi(req *domain.Absensi) (*domain.Absensi, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Absensi{}, err
	}

	createdAbsensi := &domain.Absensi{}
	err = a.DB.
		Last(createdAbsensi).
		Error

	if err != nil {
		return &domain.Absensi{}, err
	}

	return createdAbsensi, nil
}
