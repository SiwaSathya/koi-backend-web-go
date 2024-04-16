package repository

import (
	"errors"
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgrePembayaranRepository struct {
	DB *gorm.DB
}

func NewPostgrePembayaran(client *gorm.DB) domain.PembayaranRepository {
	return &posgrePembayaranRepository{
		DB: client,
	}
}

func (a *posgrePembayaranRepository) CreatePembayaran(req *domain.Pembayaran) (*domain.Pembayaran, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Pembayaran{}, err
	}

	createdPembayaran := &domain.Pembayaran{}
	err = a.DB.
		Last(createdPembayaran).
		Error

	if err != nil {
		return &domain.Pembayaran{}, err
	}

	return createdPembayaran, nil
}

func (a *posgrePembayaranRepository) GetEventByMahasiswaID(id uint) ([]domain.Pembayaran, error) {
	var res []domain.Pembayaran
	err := a.DB.
		Model(domain.Pembayaran{}).
		Where("mahasiswa_id = ?", id).
		Preload("Event").
		Find(&res).Error
	if err != nil {
		return []domain.Pembayaran{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Pembayaran{}, fmt.Errorf("record not found")
	}

	return res, nil

}
