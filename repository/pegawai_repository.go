package repository

import (
	"context"
	"database/sql"
	"pegawaiServices/model/domain"
)

type PegawaiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error)
	Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Pegawai, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error)
	FindByNip(ctx context.Context, tx *sql.Tx, nip string) (domain.Pegawai, error)
	FindByKodeOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) (domain.Opd, error)
}
