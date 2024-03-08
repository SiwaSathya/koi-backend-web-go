package usecase

import (
	"context"
	"fmt"
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"
	"koi-backend-web-go/utils"
	"time"
)

type userUseCase struct {
	userRepository      domain.UserRepository
	ormawaRepository    domain.OrmawaRepository
	mahasiswaRepository domain.MahasiswaRepository
	contextTimeout      time.Duration
}

func NewLocationUseCase(usr domain.UserRepository, orm domain.OrmawaRepository, mhs domain.MahasiswaRepository, t time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository:      usr,
		ormawaRepository:    orm,
		mahasiswaRepository: mhs,
		contextTimeout:      t,
	}
}

func (c *userUseCase) CreateUser(ctx context.Context, req *domain.CreateUser) (*domain.User, error) {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}

	pay := domain.User{
		Username: req.Username,
		Password: password,
		Role:     req.Role,
	}

	if req.Role == "ormawa" {
		pay.Username = req.NamaOrmawa
	} else {
		pay.Username = req.Username
	}

	fmt.Println(pay.Username)
	_, err = c.userRepository.GetUser(*pay.Username)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password is not match")
	}

	res, err := c.userRepository.CreateUser(&pay)
	if err != nil {
		return nil, err
	}

	if req.Role == "ormawa" {
		payOrm := domain.Ormawa{
			Password: password,
			UserID:   res.ID,
		}

		if req.NamaOrmawa != nil {
			payOrm.NamaOrmawa = *req.NamaOrmawa
			payOrm.Status = *req.Status
		} else {
			return nil, fmt.Errorf("nama ormawa and status must be field if the role is ormawa")
		}
		_, err := c.ormawaRepository.CreateOrmawa(&payOrm)
		if err != nil {
			return nil, err
		}
	} else if req.Role == "mahasiswa" {
		payMhs := domain.Mahasiswa{
			Username: *req.Username,
			Password: password,
			UserID:   res.ID,
		}
		if req.Nim != nil {
			payMhs.Nim = *req.Nim
		} else {
			return nil, fmt.Errorf("nim must be field if the role is mahasiswa")
		}
		_, err := c.mahasiswaRepository.CreateMahasiswa(&payMhs)
		if err != nil {
			return nil, err
		}
	}

	return &pay, nil
}

func (c *userUseCase) LoginUser(ctx context.Context, req *domain.LoginPayload) (*domain.User, string, error) {
	res, err := c.userRepository.GetUser(*req.Username)
	if err != nil {
		return nil, "", err
	}
	err = utils.VerifyPassword(req.Password, res.Password)
	if err != nil {
		return nil, "", fmt.Errorf("error verifying password: %v", err)
	}
	tokPay := domain.TokenClaims{
		User: res,
	}
	token, err := middleware.CreateToken(&tokPay)
	if err != nil {
		return nil, "", fmt.Errorf("cannot create token: %v", err)
	}
	return res, token, nil
}
