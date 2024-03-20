package repository

import (
	"errors"
	"fmt"
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

func (a *posgreDetailKegiatanRepository) GetDetailKegiatanByID(id uint) (*domain.DetailKegiatan, error) {
	var res domain.DetailKegiatan
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Preload("Event").
		Preload("Event.Ormawa").
		Preload("MetodePembayaran").
		Preload("Narahubung").
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.DetailKegiatan{}, fmt.Errorf("record not found")
	}

	fmt.Println(res)

	return &res, nil
}
