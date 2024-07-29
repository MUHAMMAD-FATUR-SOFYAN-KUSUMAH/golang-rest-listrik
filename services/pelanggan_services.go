package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type PelangganServices interface {
	FindAll(ctx context.Context) []response.PelangganResponse
	FindById(ctx context.Context, id int) response.PelangganResponse
	Save(ctx context.Context, pelanggan request.AddPelanggan)
	Update(ctx context.Context, pelanggan request.UpdatePelanggan)
	Delete(ctx context.Context, pelanggan request.PelangganRequest)
}
