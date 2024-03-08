package repository

import (
	"errors"
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreUserRepository struct {
	DB *gorm.DB
}

func NewPostgreUser(client *gorm.DB) domain.UserRepository {
	return &posgreUserRepository{
		DB: client,
	}
}

func (a *posgreUserRepository) CreateUser(req *domain.User) (*domain.User, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.User{}, err
	}

	return req, nil
}

func (a *posgreUserRepository) GetUser(username string) (*domain.User, error) {
	var res domain.User
	err := a.DB.
		Model(domain.User{}).
		Preload("Ormawa").
		Preload("Mahasiswa").
		Where("username = ?", username).
		Take(&res).Error
	if err != nil {
		return &domain.User{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.User{}, fmt.Errorf("record not found")
	}

	fmt.Println(res)

	return &res, nil
}
