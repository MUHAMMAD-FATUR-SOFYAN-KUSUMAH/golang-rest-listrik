package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type PembayaranServices interface {
	FindAllKonfrimasi(ctx context.Context) []response.PembayaranConfResponse
	Delete(ctx context.Context, pembayaran request.LevelSearch)
	FindAllDetails(ctx context.Context) []response.PembayaranDetailResponse
}
