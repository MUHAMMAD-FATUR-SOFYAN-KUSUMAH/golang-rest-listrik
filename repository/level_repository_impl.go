package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type LevelRepositoryImpl struct{}

// Delete implements LevelRepository.
func (repository *LevelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool) {
	sql := "DELETE FROM level WHERE id_level = $1"
	_, err := tx.ExecContext(ctx, sql, level.Id_level)
	helper.Err(err)

	defer func() {
		done <- true
	}()
}

// FindAll implements LevelRepository.
func (*LevelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Level) {
	temp := []domain.Level{}
	sql := "SELECT * FROM level"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	for row.Next() {
		var gory domain.Level
		err := row.Scan(&gory.Id_level, &gory.Nama_level)
		helper.Err(err)
		temp = append(temp, gory)
	}

	defer func() {
		done <- temp
	}()
}

// FindById implements LevelRepository.
func (*LevelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Level) {
	sql := "SELECT * FROM level WHERE id_level = $1"
	row := tx.QueryRowContext(ctx, sql, id)
	temp := domain.Level{}
	err := row.Scan(&temp.Id_level, &temp.Nama_level)
	helper.Err(err)
	fmt.Println(temp)

	defer func() {
		chann <- temp
	}()
}

// Save implements LevelRepository.
func (controller *LevelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool) {
	sql := "INSERT INTO level (nama_level) VALUES ($1) RETURNING id_level"
	err := tx.QueryRowContext(ctx, sql, level.Nama_level).Scan(&level.Id_level)
	helper.Err(err)
	defer func() {
		done <- true
	}()
}

// Update implements LevelRepository.
func (*LevelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, level domain.Level, done chan bool) {
	sql := "UPDATE level SET nama_level = $1 WHERE id_level = $2"
	_, err := tx.ExecContext(ctx, sql, level.Nama_level, level.Id_level)
	helper.Err(err)

	defer func() {
		done <- true
	}()
}

func NewLevelRepository() LevelRepository {
	return &LevelRepositoryImpl{}
}
