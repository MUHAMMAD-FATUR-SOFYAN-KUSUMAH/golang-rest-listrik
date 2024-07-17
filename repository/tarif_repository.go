package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type TarifRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Tarifkwh
	Save(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh)
	Update(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh)
	Delete(ctx context.Context, tx *sql.Tx, tarif domain.Tarifkwh)
	FindById(ctx context.Context, tx *sql.Tx, id int32) (domain.Tarifkwh, error)
}
