package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type PenggunaanServices interface {
	FindAll(ctx context.Context) []response.PenggunaanResponseIndex
	FindAllDetail(ctx context.Context, id int) []response.PenggunaanResponseDetail
	FindById(ctx context.Context, id int, chann chan response.PenggunaanResponseDetail)
	Save(ctx context.Context, penggunaan request.PenggunaanAdd)
	Update(ctx context.Context, penggunaan request.PenggunaanUpdate)
	Delete(ctx context.Context, penggunaan request.PenggunaanSearch, chann chan bool)
}
