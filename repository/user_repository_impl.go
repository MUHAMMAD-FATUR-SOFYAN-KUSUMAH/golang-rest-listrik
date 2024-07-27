package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type UserRepositoryImpl struct{}

// Delete implements UserRepository.
func (*UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool) {
	sql := "DELETE FROM public.user WHERE id_user = $1"
	_, err := tx.ExecContext(ctx, sql, user.Id_user)
	helper.Err(err)

	defer func() {
		done <- true
	}()
}

// FindAll implements UserRepository.
func (*UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.User) {
	sql := "select u.id_user, u.username, u.nama_admin, u.level, l.nama_level, l.id_level from public.user as u inner join public.level as l on u.level = l.id_level"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	defer row.Close()

	var model_user []domain.User
	for row.Next() {
		var gory domain.User
		err := row.Scan(&gory.Id_user, &gory.Username, &gory.Name, &gory.Role_level.Id_level, &gory.Role_level.Nama_level, &gory.Level)
		helper.Err(err)
		model_user = append(model_user, gory)
	}

	defer func() {
		done <- model_user
	}()
}

// FindById implements UserRepository.
func (*UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.User) {
	temp := domain.User{}
	sql := "select u.id_user, u.username, u.nama_admin, u.level, l.nama_level, l.id_level from public.user as u inner join public.level as l on u.level = l.id_level where u.id_user = $1"
	err := tx.QueryRowContext(ctx, sql, id).Scan(&temp.Id_user, &temp.Username, &temp.Name, &temp.Role_level.Id_level, &temp.Role_level.Nama_level, &temp.Level)

	helper.Err(err)
	defer func() {
		chann <- temp
	}()
}

// Save implements UserRepository.
func (*UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool) {
	sql := "INSERT INTO public.user (username, password, nama_admin, level) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, sql, user.Username, user.Password, user.Name, user.Level)
	helper.Err(err)
	defer func() {
		done <- true
	}()
}

// Update implements UserRepository.
func (*UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool) {
	sql := "UPDATE public.user SET username=$1, password=$2, nama_admin=$3, level=$4 WHERE id_user=$5"
	_, err := tx.ExecContext(ctx, sql, user.Username, user.Password, user.Name, user.Level, user.Id_user)
	helper.Err(err)
	defer func() {
		done <- true
	}()
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
