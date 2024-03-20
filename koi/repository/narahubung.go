package repository

import (
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreNarahubungRepository struct {
	DB *gorm.DB
}

func NewPostgreNarahubung(client *gorm.DB) domain.NarahubungRepository {
	return &posgreNarahubungRepository{
		DB: client,
	}
}

func (a *posgreNarahubungRepository) CreateNarahubung(req *domain.Narahubung) (*domain.Narahubung, error) {
	fmt.Println(req)
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Narahubung{}, err
	}

	createdNarahubung := &domain.Narahubung{}
	err = a.DB.
		Last(createdNarahubung).
		Error

	if err != nil {
		return &domain.Narahubung{}, err
	}

	return createdNarahubung, nil
}
