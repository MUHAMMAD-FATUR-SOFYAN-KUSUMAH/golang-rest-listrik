package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type PembayaranRepository interface {
	FindAllKonfrimasi(ctx context.Context, tx *sql.Tx, done chan []domain.Pembayaran)
	Delete(ctx context.Context, tx *sql.Tx, pembayaran domain.Pembayaran)
	FindAllDetails(ctx context.Context, tx *sql.Tx, done chan []domain.Pembayaran)
}
