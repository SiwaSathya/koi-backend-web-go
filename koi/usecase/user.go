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
	eventRepository     domain.EventRepository
	contextTimeout      time.Duration
}

func NewUserUseCase(usr domain.UserRepository, orm domain.OrmawaRepository, mhs domain.MahasiswaRepository, er domain.EventRepository, t time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository:      usr,
		ormawaRepository:    orm,
		mahasiswaRepository: mhs,
		eventRepository:     er,
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

	// fmt.Println(pay.Username)
	_, err = c.userRepository.GetUser(*pay.Username)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password is not match")
	}

	// fmt.Println("PAY", pay)
	res, err := c.userRepository.CreateUser(&pay)
	if err != nil {
		return nil, err
	}

	if req.Role == "ormawa" {
		payOrm := domain.Ormawa{
			UserID: res.ID,
		}

		if req.NamaOrmawa != nil {
			payOrm.NamaOrmawa = *req.NamaOrmawa
			payOrm.Status = *req.Status
			payOrm.Deskripsi = req.Deskripsi
			payOrm.JenisOrmawa = req.JenisOrmawa
		} else {
			return nil, fmt.Errorf("nama_ormawa and status must be field if the role is ormawa")
		}
		_, err := c.ormawaRepository.CreateOrmawa(&payOrm)
		if err != nil {
			return nil, err
		}
	} else if req.Role == "mahasiswa" {
		payMhs := domain.Mahasiswa{
			UserID: res.ID,
		}
		if req.Nim != nil {
			payMhs.NamaMahasiswa = req.NamaMahasiswa
			payMhs.Nim = *req.Nim
			payMhs.NoTelepon = req.NoTelepon
			payMhs.Email = req.Email
			payMhs.TanggalLahir = req.TanggalLahir
			payMhs.JenisKelamin = req.JenisKelamin
			payMhs.TempatLahir = req.TempatLahir
			payMhs.AlamatTinggal = req.AlamatTinggal
		} else {
			return nil, fmt.Errorf("nim must be field in if the role is mahasiswa")
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

func (c *userUseCase) GetUserById(ctx context.Context, id uint) (*domain.User, error) {
	return c.userRepository.GetUserById(id)
}

func (c *userUseCase) PengajuanEventOrmawa(ctx context.Context) (map[string]any, error) {
	res, err := c.eventRepository.GetAllEvents()
	if err != nil {
		return nil, err
	}
	var (
		waitingValidation int
		eventAccepted     int
		seminar           int
		race              int
		entertainment     int
		workshop          int
		socialActivities  int
	)
	for _, val := range res {
		if val.Category == "seminar" {
			seminar += 1
		} else if val.Category == "race" {
			race += 1
		} else if val.Category == "entertainment" {
			entertainment += 1
		} else if val.Category == "workshop" {
			workshop += 1
		} else if val.Category == "social_activies" {
			socialActivities += 1
		}

		if val.DetailKegiatan.Status == "pending" {
			waitingValidation += 1
		} else if val.DetailKegiatan.Status == "accepted" {
			eventAccepted += 1
		}

	}
	response := map[string]any{
		"event":              len(res),
		"waiting_validation": waitingValidation,
		"event_accepted":     eventAccepted,
		"seminar":            seminar,
		"race":               race,
		"entertainment":      entertainment,
		"workshop":           workshop,
		"social_activities":  socialActivities,
	}
	return response, nil
}

func (c *userUseCase) ResetPassword(ctx context.Context, req *domain.ResetPassword) error {
	if req.Password != req.PasswordConfirmation {
		return fmt.Errorf("passwords does not match")
	}
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("unable to hash password: %v", err)
	}
	req.Password = password
	err = c.userRepository.UpdatePassword(req)
	if err != nil {
		return err
	}
	return nil

}

func (c *userUseCase) UpdateProfile(ctx context.Context, req *domain.CreateUser) error {
	if req.Role == "ormawa" {
		payOrm := domain.Ormawa{
			UserID: req.ID,
		}

		if req.NamaOrmawa != nil {
			payOrm.NamaOrmawa = *req.NamaOrmawa
			payOrm.Status = *req.Status
			payOrm.Email = req.Email
			payOrm.Deskripsi = req.Deskripsi
			payOrm.JenisOrmawa = req.JenisOrmawa
		} else {
			return fmt.Errorf("nama_ormawa and status must be field if the role is ormawa")
		}
		err := c.ormawaRepository.Updateormawa(&payOrm)
		if err != nil {
			return err
		}
	} else if req.Role == "mahasiswa" {
		payMhs := domain.Mahasiswa{
			UserID: req.ID,
		}
		fmt.Println("Nama Mahasiswa", *req.NamaMahasiswa)
		if req.Nim != nil {
			payMhs.NamaMahasiswa = req.NamaMahasiswa
			payMhs.Nim = *req.Nim
			payMhs.NoTelepon = req.NoTelepon
			payMhs.Email = req.Email
			payMhs.TanggalLahir = req.TanggalLahir
			payMhs.JenisKelamin = req.JenisKelamin
			payMhs.TempatLahir = req.TempatLahir
			payMhs.AlamatTinggal = req.AlamatTinggal
		} else {
			return fmt.Errorf("nim must be field in if the role is mahasiswa")
		}
		fmt.Println(payMhs)
		err := c.mahasiswaRepository.UpdateMahasiswa(&payMhs)
		if err != nil {
			return err
		}
	}

	return nil
}
