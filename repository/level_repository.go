package repository

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
)

type LevelRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Level)
	Save(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool)
	Update(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool)
	Delete(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool)
	FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Level)
}
