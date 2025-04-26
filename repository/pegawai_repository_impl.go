package repository

import (
	"context"
	"database/sql"
	"pegawaiServices/helper"
	"pegawaiServices/model/domain"
)

type PegawaiRepositoryImpl struct {
}

func NewPegawaiRepositoryImpl() *PegawaiRepositoryImpl {
	return &PegawaiRepositoryImpl{}
}

func (repository *PegawaiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) (domain.Pegawai, error) {
	script := "INSERT INTO tb_pegawai (id, nama, nip, kode_opd) VALUES (?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, pegawai.Id, pegawai.Nama, pegawai.Nip, pegawai.KodeOpd)
	if err != nil {
		return domain.Pegawai{}, err
	}
	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pegawai domain.Pegawai) domain.Pegawai {
	script := "UPDATE tb_pegawai SET nama = ?, nip = ?, kode_opd = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, pegawai.Nama, pegawai.Nip, pegawai.KodeOpd, pegawai.Id)
	helper.PanicIfError(err)
	return pegawai
}

func (repository *PegawaiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	script := "UPDATE tb_pegawai SET deleted_at = NOW() WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)
	return nil
}

func (repository *PegawaiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Pegawai, error) {
	script := "SELECT id, nama, nip, kode_opd FROM tb_pegawai WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	pegawai := domain.Pegawai{}
	if rows.Next() {
		err = rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.KodeOpd)
		helper.PanicIfError(err)
	}
	return pegawai, nil
}

func (repository *PegawaiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.Pegawai, error) {
	script := "SELECT id, nama, nip, kode_opd FROM tb_pegawai WHERE kode_opd = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script, kodeOpd)
	helper.PanicIfError(err)
	defer rows.Close()

	pegawais := []domain.Pegawai{}
	for rows.Next() {
		pegawai := domain.Pegawai{}
		err = rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.KodeOpd)
		helper.PanicIfError(err)
		pegawais = append(pegawais, pegawai)
	}
	return pegawais, nil
}

func (repository *PegawaiRepositoryImpl) FindByNip(ctx context.Context, tx *sql.Tx, nip string) (domain.Pegawai, error) {
	script := "SELECT id, nama, nip, kode_opd FROM tb_pegawai WHERE nip = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script, nip)
	helper.PanicIfError(err)
	defer rows.Close()

	pegawai := domain.Pegawai{}
	if rows.Next() {
		err = rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Nip, &pegawai.KodeOpd)
		helper.PanicIfError(err)
	}
	return pegawai, nil
}
