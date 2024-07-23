package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type TarifRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Tarif)
	Save(ctx context.Context, tx *sql.Tx, tarif domain.Tarif)
	Update(ctx context.Context, tx *sql.Tx, tarif domain.Tarif)
	Delete(ctx context.Context, tx *sql.Tx, tarif domain.Tarif)
	FindById(ctx context.Context, tx *sql.Tx, id string, chann chan domain.Tarif)
}
