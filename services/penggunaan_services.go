package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type PenggunaanServices interface {
	FindAll(ctx context.Context, id int) []response.PenggunaanResponseDetail
	FindById(ctx context.Context, id int, chann chan response.PenggunaanResponseDetail)
	Save(ctx context.Context, penggunaan request.PenggunaanSave)
	Update(ctx context.Context, penggunaan request.PenggunaanUpdated)
	Delete(ctx context.Context, penggunaan request.PenggunaanSearch)
}
