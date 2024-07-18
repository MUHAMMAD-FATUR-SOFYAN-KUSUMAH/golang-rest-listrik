package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type TarifServices interface {
	FIndAll(ctx context.Context) []response.TarifResponse
	FindById(ctx context.Context, id request.TarifSearch) response.TarifResponse
	Save(ctx context.Context, tarif request.TarifSave)
	Update(ctx context.Context, tarif request.TarifUpdated)
	Delete(ctx context.Context, tarif request.TarifSearch)
}
