package repository

import (
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreEventRepository struct {
	DB *gorm.DB
}

func NewPostgreEvent(client *gorm.DB) domain.EventRepository {
	return &posgreEventRepository{
		DB: client,
	}
}

func (a *posgreEventRepository) CreateEvent(req *domain.Event) (*domain.Event, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.Event{}, err
	}

	createdEvent := &domain.Event{}
	err = a.DB.
		Last(createdEvent).
		Error

	if err != nil {
		return &domain.Event{}, err
	}

	return createdEvent, nil
}
