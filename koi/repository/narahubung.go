package repository

import (
	"errors"
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

func (a *posgreNarahubungRepository) UpdateNarahubung(req *domain.Narahubung) error {
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Where("id = ?", req.ID).
		Select("judul", "nama_narahubung", "no_telepon").
		Updates(map[string]interface{}{
			"judul":      req.Judul,
			"nama_bank":  req.NamaNarahubung,
			"no_telepon": req.NoTelepon,
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
