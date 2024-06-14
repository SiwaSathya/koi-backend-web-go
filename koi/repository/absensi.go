package repository

import (
	"errors"
	"fmt"
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

func (a *posgreAbsensiRepository) GetAbsensiByEventID(eventId uint) ([]domain.Absensi, error) {
	var res []domain.Absensi
	err := a.DB.
		Model(domain.Absensi{}).
		Where("event_id = ?", eventId).
		// Preload("User").
		Find(&res).Error
	if err != nil {
		return []domain.Absensi{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Absensi{}, fmt.Errorf("record not found")
	}

	return res, nil
}

func (a *posgreAbsensiRepository) UpdateStatus(eventID uint, userID uint, status string) error {
	fmt.Println(eventID, userID, status)
	err := a.DB.
		Model(domain.Absensi{}).
		Where("event_id = ?", eventID).
		Where("user_id = ?", userID).
		Select("status").
		Updates(map[string]interface{}{
			"status": status,
		}).
		Error

	if err != nil {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("record not found")
	}

	return nil
}

func (a *posgreAbsensiRepository) GetAllAbsensi() ([]domain.Absensi, error) {
	var res []domain.Absensi
	err := a.DB.
		Model(domain.Absensi{}).
		Preload("Event").
		Preload("Event.Pembayaran").
		Find(&res).Error
	if err != nil {
		return []domain.Absensi{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Absensi{}, fmt.Errorf("record not found")
	}

	return res, nil
}
