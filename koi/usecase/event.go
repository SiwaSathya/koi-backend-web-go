package usecase

import (
	"context"
	"fmt"
	"koi-backend-web-go/domain"
	"time"

	"github.com/pandeptwidyaop/golog"
)

type eventUseCase struct {
	detailKegiatanRepository  domain.DetailKegiatanRepository
	eventRepository           domain.EventRepository
	narahubungRepository      domain.NarahubungRepository
	metodePembayranRepository domain.MetodePembayaranRepository
	userRepository            domain.UserRepository
	contextTimeout            time.Duration
}

func NewEventUseCase(dtl domain.DetailKegiatanRepository, evt domain.EventRepository, nrh domain.NarahubungRepository, mtd domain.MetodePembayaranRepository, usr domain.UserRepository, t time.Duration) domain.EventUseCase {
	return &eventUseCase{
		detailKegiatanRepository:  dtl,
		eventRepository:           evt,
		narahubungRepository:      nrh,
		metodePembayranRepository: mtd,
		userRepository:            usr,
		contextTimeout:            t,
	}
}

func (e *eventUseCase) CreateEvent(ctx context.Context, req *domain.CreateEvent) (*domain.Event, error) {
	id, err := e.userRepository.GetUserById(req.OrmawaID)
	if err != nil {
		return nil, err
	}
	payEvent := domain.Event{
		OrmawaID:        id.Ormawa.ID,
		NamaKegiatan:    req.NamaKegiatan,
		TanggalKegiatan: req.TanggalKegiatan,
		TingkatKegiatan: req.TingkatKegiatan,
		HargaTiket:      req.HargaTiket,
	}
	res, err := e.eventRepository.CreateEvent(&payEvent)
	if err != nil {
		return nil, err
	}

	req.DetailKegiatan.EventID = res.ID
	resDet, err := e.detailKegiatanRepository.CreateDetailKegiatan(&req.DetailKegiatan)
	if err != nil {
		return nil, err
	}

	for _, valMet := range req.MetodePembayaran {
		valMet.DetailKegiatanID = resDet.ID
		_, err := e.metodePembayranRepository.CreateMetodePembayaran(&valMet)
		if err != nil {
			golog.Slack.Error(fmt.Sprintf("cannot store the metode pembayaran %v", valMet), err)
			continue
		}
	}

	for _, valNar := range req.Narahubung {
		valNar.DetailKegiatanID = resDet.ID
		_, err := e.narahubungRepository.CreateNarahubung(&valNar)
		if err != nil {
			golog.Slack.Error(fmt.Sprintf("cannot store the narahubung %v", valNar), err)
			continue
		}
	}

	return res, nil
}
