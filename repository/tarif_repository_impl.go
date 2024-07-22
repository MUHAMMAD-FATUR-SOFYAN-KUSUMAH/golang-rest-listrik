package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type TarifRepositoryImpl struct{}

// Delete implements TarifRepository.
func (tf *TarifRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tarif domain.Tarif) {
	sql := "DELETE FROM tarif WHERE uuid_tarif = $1"
	_, err := tx.ExecContext(ctx, sql, tarif.UuidTarif)
	helper.Err(err)
}

// FindAll implements TarifRepository.
func (tf *TarifRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Tarif) {
	sql := "SELECT uuid_tarif, daya, tarifperkwh FROM tarif"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	defer row.Close()

	var model_tarif []domain.Tarif
	defer func() {
		done <- model_tarif
	}()

	for row.Next() {
		var gory domain.Tarif
		err := row.Scan(&gory.UuidTarif, &gory.Daya, &gory.TarifPerKwh)
		helper.Err(err)
		model_tarif = append(model_tarif, gory)
	}
}

// FindById implements TarifRepository.
func (tf *TarifRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int32) (domain.Tarif, error) {
	panic("unimplemented")
}

// Save implements TarifRepository.
func (tf *TarifRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, tarif domain.Tarif) {
	uuid := helper.GenerateRandomToken()
	sql := "INSERT INTO tarif (daya, tarifperkwh, uuid_tarif) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, sql, tarif.Daya, tarif.TarifPerKwh, uuid)
	helper.Err(err)
}

// Update implements TarifRepository.
func (tf *TarifRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tarif domain.Tarif) {
	sql := "UPDATE tarif SET daya=$1, tarifperkwh=$2 where uuid_tarif=$3"
	_, err := tx.ExecContext(ctx, sql, tarif.Daya, tarif.TarifPerKwh, tarif.UuidTarif)
	helper.Err(err)
}

func NewTarifRepository() TarifRepository {
	return &TarifRepositoryImpl{}
}
