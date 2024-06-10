package repository

import (
	"errors"
	"fmt"
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

func (a *posgreEventRepository) GetAllEvents() ([]domain.Event, error) {
	var res []domain.Event
	err := a.DB.
		Model(domain.Event{}).
		Preload("DetailKegiatan").
		Find(&res).Error
	if err != nil {
		return []domain.Event{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Event{}, fmt.Errorf("record not found")
	}

	return res, nil
}
func (a *posgreEventRepository) GetEventByID(id uint) (*domain.Event, error) {
	var res domain.Event
	err := a.DB.
		Model(domain.Event{}).
		Where("id = ?", id).
		Preload("DetailKegiatan").
		Preload("Narahubung").
		Preload("MetodePembayaran").
		Find(&res).Error
	if err != nil {
		return &domain.Event{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Event{}, fmt.Errorf("record not found")
	}

	return &res, nil
}

func (a *posgreEventRepository) GetEventByOrmawaID(id uint) ([]domain.Event, error) {
	var res []domain.Event
	err := a.DB.
		Model(domain.Event{}).
		Where("ormawa_id = ?", id).
		Preload("DetailKegiatan").
		Find(&res).Error
	if err != nil {
		return []domain.Event{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Event{}, fmt.Errorf("record not found")
	}

	return res, nil
}

func (a *posgreEventRepository) UpdateEvent(req *domain.Event) error {
	err := a.DB.
		Model(domain.Event{}).
		Where("id = ?", req.ID).
		Select("nama_kegiatan", "harga_tiket", "its_open", "categry", "tanggal_kegiatan", "tingkat_kegiatan").
		Updates(map[string]interface{}{
			"nama_kegiatan":    req.NamaKegiatan,
			"harga_tiket":      req.HargaTiket,
			"its_open":         req.ItsOpen,
			"category":         req.Category,
			"tanggal_kegiatan": req.TanggalKegiatan,
			"tingkat_kegiatan": req.TanggalKegiatan,
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
