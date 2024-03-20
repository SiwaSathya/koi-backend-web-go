package usecase

import (
	"context"
	"koi-backend-web-go/domain"
	"time"
)

type detailKegiatanUseCase struct {
	detailKegiatanRepository domain.DetailKegiatanRepository
	contextTimeout           time.Duration
}

func NewDetailKegiatanUseCase(dtl domain.DetailKegiatanRepository, t time.Duration) domain.DetailKegiatanUseCase {
	return &detailKegiatanUseCase{
		detailKegiatanRepository: dtl,
		contextTimeout:           t,
	}
}

func (k *detailKegiatanUseCase) GetDetailKegiatanByID(ctx context.Context, id uint) (*domain.DetailKegiatan, error) {
	return k.detailKegiatanRepository.GetDetailKegiatanByID(id)
}
