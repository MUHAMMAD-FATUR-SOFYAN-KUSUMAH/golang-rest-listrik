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
	panic("unimplemented")
}

// FindAll implements TarifRepository.
func (tf *TarifRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Tarif {
	sql := "SELECT uuid_tarif, daya, tarifperkwh FROM tarif"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)
	defer row.Close()

	var model_tarif []domain.Tarif
	for row.Next() {
		var gory domain.Tarif
		err := row.Scan(&gory.UuidTarif, &gory.Daya, &gory.TarifPerKwh)
		helper.Err(err)
		model_tarif = append(model_tarif, gory)
	}

	return model_tarif
}

// FindById implements TarifRepository.
func (tf *TarifRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int32) (domain.Tarif, error) {
	panic("unimplemented")
}

// Save implements TarifRepository.
func (tf *TarifRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, tarif domain.Tarif) {
	uuid := helper.GenerateRandomToken()
	sql := "INSERT INTO (daya, tarifperkwh, uuid_tarif) (?, ?, ?)"
	tx.QueryContext(ctx, sql, tarif.Daya, tarif.TarifPerKwh, uuid)
}

// Update implements TarifRepository.
func (tf *TarifRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tarif domain.Tarif) {
	panic("unimplemented")
}

func NewTarifRepository() TarifRepository {
	return &TarifRepositoryImpl{}
}
