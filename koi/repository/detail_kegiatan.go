package repository

import (
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreDetailKegiatanRepository struct {
	DB *gorm.DB
}

func NewPostgreDetailKegiatan(client *gorm.DB) domain.DetailKegiatanRepository {
	return &posgreDetailKegiatanRepository{
		DB: client,
	}
}

func (a *posgreDetailKegiatanRepository) CreateDetailKegiatan(req *domain.DetailKegiatan) (*domain.DetailKegiatan, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	createdDetailKegiatan := &domain.DetailKegiatan{}
	err = a.DB.
		Last(createdDetailKegiatan).
		Error

	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	return createdDetailKegiatan, nil
}
