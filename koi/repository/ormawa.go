package repository

import (
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreOrmawaRepository struct {
	DB *gorm.DB
}

func NewPostgreOrmawa(client *gorm.DB) domain.OrmawaRepository {
	return &posgreOrmawaRepository{
		DB: client,
	}
}

func (a *posgreOrmawaRepository) CreateOrmawa(req *domain.Ormawa) (*domain.Ormawa, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Ormawa{}, err
	}

	return req, nil
}
