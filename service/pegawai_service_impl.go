package service

import (
	"context"
	"database/sql"
	"fmt"
	"pegawaiServices/helper"
	"pegawaiServices/model/domain"
	"pegawaiServices/model/web"
	"pegawaiServices/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PegawaiServiceImpl struct {
	pegawaiRepository repository.PegawaiRepository
	db                *sql.DB
	validate          *validator.Validate
}

func NewPegawaiServiceImpl(pegawaiRepository repository.PegawaiRepository, db *sql.DB, validate *validator.Validate) *PegawaiServiceImpl {
	return &PegawaiServiceImpl{
		pegawaiRepository: pegawaiRepository,
		db:                db,
		validate:          validate,
	}
}

func (service *PegawaiServiceImpl) Create(ctx context.Context, request web.PegawaiCreateRequest) (web.PegawaiResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	ts := time.Now().Format("060102150405")
	currentYear := time.Now().Year()
	u := uuid.New()
	randomPart := u.String()[0:6]
	id := fmt.Sprintf("PEG-%v-%v-%v", currentYear, ts, randomPart)

	request.Id = id

	pegawai, err := service.pegawaiRepository.Create(ctx, tx, domain.Pegawai{
		Id:      request.Id,
		Nama:    request.Nama,
		Nip:     request.Nip,
		KodeOpd: request.KodeOpd,
	})
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		KodeOpd: pegawai.KodeOpd,
	}, nil
}

func (service *PegawaiServiceImpl) Update(ctx context.Context, request web.PegawaiUpdateRequest) (web.PegawaiResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawai := service.pegawaiRepository.Update(ctx, tx, domain.Pegawai{
		Id:      request.Id,
		Nama:    request.Nama,
		Nip:     request.Nip,
		KodeOpd: request.KodeOpd,
	})

	return web.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		KodeOpd: pegawai.KodeOpd,
	}, nil
}

func (service *PegawaiServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.pegawaiRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *PegawaiServiceImpl) FindById(ctx context.Context, id string) (web.PegawaiResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawai, err := service.pegawaiRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		KodeOpd: pegawai.KodeOpd,
	}, nil
}

func (service *PegawaiServiceImpl) FindAll(ctx context.Context, kodeOpd string) ([]web.PegawaiResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawais, err := service.pegawaiRepository.FindAll(ctx, tx, kodeOpd)
	if err != nil {
		return []web.PegawaiResponse{}, err
	}

	pegawaiResponses := []web.PegawaiResponse{}
	for _, pegawai := range pegawais {
		pegawaiResponses = append(pegawaiResponses, web.PegawaiResponse{
			Id:      pegawai.Id,
			Nama:    pegawai.Nama,
			Nip:     pegawai.Nip,
			KodeOpd: pegawai.KodeOpd,
		})
	}

	return pegawaiResponses, nil
}

func (service *PegawaiServiceImpl) FindByNip(ctx context.Context, nip string) (web.PegawaiResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.PegawaiResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	pegawai, err := service.pegawaiRepository.FindByNip(ctx, tx, nip)
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	return web.PegawaiResponse{
		Id:      pegawai.Id,
		Nama:    pegawai.Nama,
		Nip:     pegawai.Nip,
		KodeOpd: pegawai.KodeOpd,
	}, nil
}
