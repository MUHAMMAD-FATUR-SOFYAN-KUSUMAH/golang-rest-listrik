package services

import (
	"context"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type UserServices interface {
	FindAll(ctx context.Context) []response.UserResponse
	Save(ctx context.Context, user request.AddUser)
	Update(ctx context.Context, user request.UserUpdate)
	Delete(ctx context.Context, user request.UserSearch)
	FindById(ctx context.Context, id int) response.UserResponse
}
