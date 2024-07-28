package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type PelangganRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Pelanggan)
	FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Pelanggan)
	Save(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan)
	Update(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan)
	Delete(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan)
}
