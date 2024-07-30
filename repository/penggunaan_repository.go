package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type PenggunaanRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Penggunaan, id int)
	Save(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan)
	Update(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan)
	Delete(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan)
	FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Penggunaan)
}
