package repository

import (
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreMetodePembayaranRepository struct {
	DB *gorm.DB
}

func NewPostgreMetodePembayaran(client *gorm.DB) domain.MetodePembayaranRepository {
	return &posgreMetodePembayaranRepository{
		DB: client,
	}
}

func (a *posgreMetodePembayaranRepository) CreateMetodePembayaran(req *domain.MetodePembayaran) (*domain.MetodePembayaran, error) {
	fmt.Println(req)
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.MetodePembayaran{}, err
	}

	createdMetodePembayaran := &domain.MetodePembayaran{}
	err = a.DB.
		Last(createdMetodePembayaran).
		Error

	if err != nil {
		return &domain.MetodePembayaran{}, err
	}

	return createdMetodePembayaran, nil
}
