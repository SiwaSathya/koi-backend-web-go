package repository

import (
	"errors"
	"fmt"
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

func (a *posgreOrmawaRepository) GetOrmawaByID(id uint) (*domain.Ormawa, error) {
	var res domain.Ormawa
	err := a.DB.
		Model(domain.Ormawa{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Ormawa{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Ormawa{}, fmt.Errorf("record not found")
	}

	fmt.Println(res)

	return &res, nil
}

func (a *posgreOrmawaRepository) Updateormawa(req *domain.Ormawa) error {
	err := a.DB.
		Model(domain.Ormawa{}).
		Where("user_id = ?", req.UserID).
		Select("nama_ormawa", "status", "deskripsi", "jenis_ormawa", "email", "cover", "logo").
		Updates(map[string]interface{}{
			"nama_ormawa":  req.NamaOrmawa,
			"status":       req.Status,
			"email":        req.Email,
			"deskripsi":    req.Deskripsi,
			"jenis_ormawa": req.JenisOrmawa,
			"cover":        req.Cover,
			"logo":         req.Logo,
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
