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
	ormawaRepository          domain.OrmawaRepository
	contextTimeout            time.Duration
}

func NewEventUseCase(dtl domain.DetailKegiatanRepository, evt domain.EventRepository, nrh domain.NarahubungRepository, mtd domain.MetodePembayaranRepository, usr domain.UserRepository, orm domain.OrmawaRepository, t time.Duration) domain.EventUseCase {
	return &eventUseCase{
		detailKegiatanRepository:  dtl,
		eventRepository:           evt,
		narahubungRepository:      nrh,
		metodePembayranRepository: mtd,
		userRepository:            usr,
		ormawaRepository:          orm,
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
		Category:        req.Category,
		TypeImplement:   req.TypeImplement,
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

func (e *eventUseCase) GetAllEvents(ctx context.Context) ([]domain.Event, error) {
	return e.eventRepository.GetAllEvents()
}

func (e *eventUseCase) GetEventByOrmawaID(ctx context.Context, id uint) (*domain.ResponeListEventOrmawa, error) {
	resOrm, err := e.ormawaRepository.GetOrmawaByID(id)
	if err != nil {
		return nil, err
	}

	resEvent, err := e.eventRepository.GetEventByOrmawaID(id)
	if err != nil {
		return nil, err
	}

	result := domain.ResponeListEventOrmawa{
		Ormawa: *resOrm,
		Event:  resEvent,
	}

	return &result, nil
}

func (e *eventUseCase) UpdateEvent(ctx context.Context, req *domain.CreateEvent) error {
	payEvent := domain.Event{
		ID:              req.ID,
		NamaKegiatan:    req.NamaKegiatan,
		TanggalKegiatan: req.TanggalKegiatan,
		TingkatKegiatan: req.TingkatKegiatan,
		HargaTiket:      req.HargaTiket,
	}
	err := e.eventRepository.UpdateEvent(&payEvent)
	if err != nil {
		return err
	}

	// req.DetailKegiatan.EventID = res.ID
	err = e.detailKegiatanRepository.UpdateDetailKegiatan(&req.DetailKegiatan)
	if err != nil {
		return err
	}

	for _, valMet := range req.MetodePembayaran {
		// valMet.DetailKegiatanID = resDet.ID
		err = e.metodePembayranRepository.UpdateMetodePembayaran(&valMet)
		if err != nil {
			golog.Slack.Error(fmt.Sprintf("cannot store the metode pembayaran %v", valMet), err)
			continue
		}
	}

	for _, valNar := range req.Narahubung {
		// valNar.DetailKegiatanID = resDet.ID
		err = e.narahubungRepository.UpdateNarahubung(&valNar)
		if err != nil {
			golog.Slack.Error(fmt.Sprintf("cannot store the narahubung %v", valNar), err)
			continue
		}
	}

	return nil
}

func (e *eventUseCase) GetEventByID(ctx context.Context, id uint) (*domain.Event, error) {
	return e.eventRepository.GetEventByID(id)
}

func (e *eventUseCase) GetEventByIDAndOrmawaID(ctx context.Context, idUser uint, idEvent uint) (*domain.Event, error) {
	fmt.Println("ini id user: ", idUser)
	fmt.Println(1)
	user, err := e.userRepository.GetUserById(idUser)
	if err != nil {
		return nil, err
	}
	fmt.Println(2)
	fmt.Println("ini user: ", user.Ormawa)
	idOrmawa := user.Ormawa.ID
	fmt.Println("ini id ormawa: ", idOrmawa)
	return e.eventRepository.GetEventByIDAndOrmawaID(idOrmawa, idEvent)
}
