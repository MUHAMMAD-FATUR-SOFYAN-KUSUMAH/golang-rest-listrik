package services

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/repository"

	"github.com/go-playground/validator/v10"
)

type UserServicesImpl struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Delete implements UserServices.
func (*UserServicesImpl) Delete(ctx context.Context, tx *sql.Tx, user request.UserSearch) {
	panic("unimplemented")
}

// FindAll implements UserServices.
func (services *UserServicesImpl) FindAll(ctx context.Context) []response.UserResponse {
	chann := make(chan []domain.User)
	tx, err := services.DB.Begin()
	helper.Err(err)

	defer close(chann)

	go services.UserRepository.FindAll(ctx, tx, chann)
	temp := <-chann
	return helper.UserResponses(temp)
}

// FindById implements UserServices.
func (*UserServicesImpl) FindById(ctx context.Context, tx *sql.Tx, id int) response.UserResponse {
	panic("unimplemented")
}

// Save implements UserServices.
func (*UserServicesImpl) Save(ctx context.Context, tx *sql.Tx, user request.AddUser) {
	panic("unimplemented")
}

// Update implements UserServices.
func (*UserServicesImpl) Update(ctx context.Context, tx *sql.Tx, user request.UserUpdate) {
	panic("unimplemented")
}

func NewUserServicesImpl(db *sql.DB, userRepository repository.UserRepository, validate *validator.Validate) UserServices {
	return &UserServicesImpl{
		DB:             db,
		UserRepository: userRepository,
		Validate:       validate,
	}
}
