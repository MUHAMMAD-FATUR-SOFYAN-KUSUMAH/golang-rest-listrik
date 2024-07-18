package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type TarifRepositoryImpl struct{}

// Delete implements TarifRepository.
func (tf *TarifRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh) {
	panic("unimplemented")
}

// FindAll implements TarifRepository.
func (tf *TarifRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Tarifkwh {
	sql := "SELECT id_tarif, daya, tarifperkwh FROM tarif"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)
	defer row.Close()

	var model_tarif []domain.Tarifkwh
	for row.Next() {
		gory := domain.Tarifkwh{}
		err := row.Scan(&gory.UuidTarif, &gory.Daya, &gory.TarifPerKwh)
		helper.Err(err)
		model_tarif = append(model_tarif, gory)
	}

	return model_tarif
}

// FindById implements TarifRepository.
func (tf *TarifRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int32) (domain.Tarifkwh, error) {
	panic("unimplemented")
}

// Save implements TarifRepository.
func (tf *TarifRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh) {
	panic("unimplemented")
}

// Update implements TarifRepository.
func (tf *TarifRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh) {
	panic("unimplemented")
}

func NewTarifRepository() TarifRepository {
	return &TarifRepositoryImpl{}
}
