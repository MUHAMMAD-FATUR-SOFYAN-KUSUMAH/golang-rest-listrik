package services

import (
	"context"
	"golang_listrik/model/request"
)

type PembayaranServices interface {
	FindAllKonfrimasi(ctx context.Context)
	Delete(ctx context.Context, pembayaran request.LevelSearch)
	FindAllDetails(ctx context.Context)
}
