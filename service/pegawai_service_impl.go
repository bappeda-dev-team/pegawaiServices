package service

import (
	"context"
	"database/sql"
	"math/rand"
	"pegawaiServices/helper"
	"pegawaiServices/model/domain"
	"pegawaiServices/model/web"
	"pegawaiServices/repository"

	"github.com/go-playground/validator/v10"
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

	statusPegawai := "valid"

	_, err = service.pegawaiRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		statusPegawai = "tidak_valid"
	}

	RandomNumber := rand.Intn(100000000000)

	pegawai, err := service.pegawaiRepository.Create(ctx, tx, domain.Pegawai{
		Id:            RandomNumber,
		Nama:          request.Nama,
		Nip:           request.Nip,
		KodeOpd:       request.KodeOpd,
		StatusPegawai: statusPegawai,
	})
	if err != nil {
		return web.PegawaiResponse{}, err
	}

	opd, _ := service.pegawaiRepository.FindByKodeOpd(ctx, tx, pegawai.KodeOpd)

	return web.PegawaiResponse{
		Nama:          pegawai.Nama,
		Nip:           pegawai.Nip,
		KodeOpd:       pegawai.KodeOpd,
		NamaOpd:       opd.NamaOpd,
		StatusPegawai: pegawai.StatusPegawai,
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

	statusPegawai := "valid"

	_, err = service.pegawaiRepository.FindByKodeOpd(ctx, tx, request.KodeOpd)
	if err != nil {
		statusPegawai = "tidak_valid"
	}

	if request.StatusPegawai != "" {
		statusPegawai = request.StatusPegawai
	}

	pegawai := service.pegawaiRepository.Update(ctx, tx, domain.Pegawai{
		Id:            request.Id,
		Nama:          request.Nama,
		Nip:           request.Nip,
		KodeOpd:       request.KodeOpd,
		StatusPegawai: statusPegawai,
	})

	opd, _ := service.pegawaiRepository.FindByKodeOpd(ctx, tx, pegawai.KodeOpd)

	return web.PegawaiResponse{
		Nama:          pegawai.Nama,
		Nip:           pegawai.Nip,
		KodeOpd:       pegawai.KodeOpd,
		NamaOpd:       opd.NamaOpd,
		StatusPegawai: pegawai.StatusPegawai,
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

	if len(pegawais) == 0 {
		return []web.PegawaiResponse{}, nil
	}

	opd, _ := service.pegawaiRepository.FindByKodeOpd(ctx, tx, pegawais[0].KodeOpd)

	pegawaiResponses := []web.PegawaiResponse{}
	for _, pegawai := range pegawais {
		pegawaiResponses = append(pegawaiResponses, web.PegawaiResponse{
			Id:            pegawai.Id,
			Nama:          pegawai.Nama,
			Nip:           pegawai.Nip,
			KodeOpd:       pegawai.KodeOpd,
			NamaOpd:       opd.NamaOpd,
			StatusPegawai: pegawai.StatusPegawai,
			CreatedAt:     pegawai.CreatedAt,
			UpdatedAt:     pegawai.UpdatedAt,
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
		Id:            pegawai.Id,
		Nama:          pegawai.Nama,
		Nip:           pegawai.Nip,
		KodeOpd:       pegawai.KodeOpd,
		StatusPegawai: pegawai.StatusPegawai,
	}, nil
}
