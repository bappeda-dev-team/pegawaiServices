package service

import (
	"context"
	"pegawaiServices/model/web"
)

type PegawaiService interface {
	Create(ctx context.Context, request web.PegawaiCreateRequest) (web.PegawaiResponse, error)
	Update(ctx context.Context, request web.PegawaiUpdateRequest) (web.PegawaiResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.PegawaiResponse, error)
	FindAll(ctx context.Context, kodeOpd string) ([]web.PegawaiResponse, error)
	FindByNip(ctx context.Context, nip string) (web.PegawaiResponse, error)
}
