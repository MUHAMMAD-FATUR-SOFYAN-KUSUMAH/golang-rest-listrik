package services

import (
	"context"
	"database/sql"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
)

type UserServices interface {
	FindAll(ctx context.Context) []response.UserResponse
	Save(ctx context.Context, tx *sql.Tx, user request.AddUser)
	Update(ctx context.Context, tx *sql.Tx, user request.UserUpdate)
	Delete(ctx context.Context, tx *sql.Tx, user request.UserSearch)
	FindById(ctx context.Context, tx *sql.Tx, id int) response.UserResponse
}
