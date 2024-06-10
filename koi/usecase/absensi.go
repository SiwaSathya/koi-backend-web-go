package usecase

import (
	"context"
	"koi-backend-web-go/domain"
	"time"
)

type absensiUseCase struct {
	absensiRepository domain.AbsensiRepoository
	contextTimeout    time.Duration
}

func NewAbsensiUseCase(dtl domain.AbsensiRepoository, t time.Duration) domain.AbsensiUseCase {
	return &absensiUseCase{
		absensiRepository: dtl,
		contextTimeout:    t,
	}
}

func (a *absensiUseCase) CreateAbsensi(ctx context.Context, req *domain.Absensi) (*domain.Absensi, error) {
	return a.absensiRepository.CreateAbsensi(req)
}
