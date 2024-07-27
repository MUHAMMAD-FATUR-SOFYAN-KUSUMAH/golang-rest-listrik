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
func (services *UserServicesImpl) Delete(ctx context.Context, user request.UserSearch) {
	err := services.Validate.Struct(user)
	helper.Err(err)

	chann := make(chan bool)
	tx, err2 := services.DB.Begin()
	helper.Err(err2)

	defer helper.Tx(tx)
	defer close(chann)

	go services.UserRepository.Delete(ctx, tx, domain.User{
		Id_user: int32(user.Id),
	}, chann)

	<-chann
}

// FindAll implements UserServices.
func (services *UserServicesImpl) FindAll(ctx context.Context) []response.UserResponse {
	chann := make(chan []domain.User)
	tx, err := services.DB.Begin()
	helper.Err(err)

	defer close(chann)

	defer helper.Tx(tx)

	go services.UserRepository.FindAll(ctx, tx, chann)
	temp := <-chann
	return helper.UserResponses(temp)
}

// FindById implements UserServices.
func (services *UserServicesImpl) FindById(ctx context.Context, id int) response.UserResponse {
	tx, _ := services.DB.Begin()
	chann := make(chan domain.User)

	defer close(chann)

	defer helper.Tx(tx)

	go services.UserRepository.FindById(ctx, tx, id, chann)

	return helper.UserReponse(<-chann)
}

// Save implements UserServices.
func (services *UserServicesImpl) Save(ctx context.Context, user request.AddUser) {
	err := services.Validate.Struct(user)
	helper.Err(err)

	userr := domain.User{
		Username: user.Username,
		Name:     user.Name,
		Password: user.Password,
		Level:    user.Role,
	}

	chann := make(chan bool)
	tx, _ := services.DB.Begin()
	defer close(chann)
	go services.UserRepository.Save(ctx, tx, userr, chann)

	<-chann
	defer helper.Tx(tx)
}

// Update implements UserServices.
func (services *UserServicesImpl) Update(ctx context.Context, user request.UserUpdate) {
	err := services.Validate.Struct(user)

	tx, _ := services.DB.Begin()
	helper.Err(err)

	makeki := make(chan bool)

	go services.UserRepository.Update(ctx, tx, domain.User{
		Id_user:  int32(user.Id_user),
		Username: user.Username,
		Name:     user.Name,
		Password: user.Password,
		Level:    user.Role,
	}, makeki)
	defer close(makeki)
	<-makeki
	defer helper.Tx(tx)
}

func NewUserServicesImpl(db *sql.DB, userRepository repository.UserRepository, validate *validator.Validate) UserServices {
	return &UserServicesImpl{
		DB:             db,
		UserRepository: userRepository,
		Validate:       validate,
	}
}
