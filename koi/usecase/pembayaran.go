package usecase

import (
	"context"
	"koi-backend-web-go/domain"
	"time"
)

type pembayaranUseCase struct {
	pembayaranRepository domain.PembayaranRepository
	mahasiswaRepository  domain.MahasiswaRepository
	contextTimeout       time.Duration
}

func NewPembayaranUseCase(dtl domain.PembayaranRepository, mhs domain.MahasiswaRepository, t time.Duration) domain.PembayaranUseCase {
	return &pembayaranUseCase{
		pembayaranRepository: dtl,
		mahasiswaRepository:  mhs,
		contextTimeout:       t,
	}
}

func (k *pembayaranUseCase) CreatePembayaran(ctx context.Context, req *domain.Pembayaran) (*domain.Pembayaran, error) {
	res, err := k.mahasiswaRepository.GetMahasiswaByUserID(req.MahasiswaID)
	if err != nil {
		return nil, err
	}
	req.MahasiswaID = res.Nim
	return k.pembayaranRepository.CreatePembayaran(req)
}

func (k *pembayaranUseCase) GetEventByMahasiswaID(ctx context.Context, id uint) ([]domain.Pembayaran, error) {
	res, err := k.mahasiswaRepository.GetMahasiswaByUserID(id)
	if err != nil {
		return nil, err
	}
	id = res.Nim
	return k.pembayaranRepository.GetEventByMahasiswaID(id)
}
