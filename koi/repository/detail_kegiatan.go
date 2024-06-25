package repository

import (
	"errors"
	"fmt"
	"koi-backend-web-go/domain"

	"gorm.io/gorm"
)

type posgreDetailKegiatanRepository struct {
	DB *gorm.DB
}

func NewPostgreDetailKegiatan(client *gorm.DB) domain.DetailKegiatanRepository {
	return &posgreDetailKegiatanRepository{
		DB: client,
	}
}

func (a *posgreDetailKegiatanRepository) CreateDetailKegiatan(req *domain.DetailKegiatan) (*domain.DetailKegiatan, error) {
	err := a.DB.
		Create(&req).
		Error

	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	createdDetailKegiatan := &domain.DetailKegiatan{}
	err = a.DB.
		Last(createdDetailKegiatan).
		Error

	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	return createdDetailKegiatan, nil
}

func (a *posgreDetailKegiatanRepository) GetDetailKegiatanByID(id uint) (*domain.DetailKegiatan, error) {
	var res domain.DetailKegiatan
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Preload("Event").
		Preload("Event.Ormawa").
		Preload("MetodePembayaran").
		Preload("Narahubung").
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.DetailKegiatan{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.DetailKegiatan{}, fmt.Errorf("record not found")
	}

	fmt.Println(res)

	return &res, nil
}

func (a *posgreDetailKegiatanRepository) UpdateStatus(eventID uint, status string) error {
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Where("event_id = ?", eventID).
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

func (a *posgreDetailKegiatanRepository) UpdateDetailKegiatan(req *domain.DetailKegiatan) error {
	err := a.DB.
		Model(domain.DetailKegiatan{}).
		Where("id = ?", req.ID).
		Select("waktu_pelaksanaan", "lokasi", "status", "deskripsi", "gambar_kegiatan", "file_pengajuan", "sertifikat").
		Updates(map[string]interface{}{
			"waktu_pelaksanaan": req.WaktuPelaksanaan,
			"lokasi":            req.Lokasi,
			"status":            req.Status,
			"deskripsi":         req.Deskripsi,
			"gambar_kegiatan":   req.GambarKegiatan,
			"file_pengajuan":    req.FilePengajuan,
			"sertifikat":        req.Sertifikat,
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
