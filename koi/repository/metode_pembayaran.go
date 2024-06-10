package repository

import (
	"errors"
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

func (a *posgreMetodePembayaranRepository) UpdateMetodePembayaran(req *domain.MetodePembayaran) error {
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Where("id = ?", req.ID).
		Select("judul", "nama_bank", "no_rekening", "pemilik_rekening").
		Updates(map[string]interface{}{
			"judul":            req.Judul,
			"nama_bank":        req.NamaBank,
			"no_rekening":      req.NoRekening,
			"pemilik_rekening": req.PemilikRekening,
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
