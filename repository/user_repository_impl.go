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
	panic("unimplemented")
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
	panic("unimplemented")
}

// Save implements UserRepository.
func (*UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (*UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User, done chan bool) {
	panic("unimplemented")
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
