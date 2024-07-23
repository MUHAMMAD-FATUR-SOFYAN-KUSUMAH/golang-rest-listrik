package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type LevelServices interface {
	FIndAll(ctx context.Context) []response.LevelResponse
	FindById(ctx context.Context, id request.LevelSearch) response.LevelResponse
	Save(ctx context.Context, level request.LevelRequest)
	Update(ctx context.Context, level request.LevelUpdate)
	Delete(ctx context.Context, level request.LevelSearch)
}
