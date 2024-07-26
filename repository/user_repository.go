package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.User)
	Save(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool)
	Update(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool)
	Delete(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool)
	FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.User)
}
