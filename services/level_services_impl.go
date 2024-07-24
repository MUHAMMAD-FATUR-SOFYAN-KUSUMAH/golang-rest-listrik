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

type LevelServicesImpl struct {
	LevelRepository repository.LevelRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

// Delete implements LevelServices.
func (sv *LevelServicesImpl) Delete(ctx context.Context, level request.LevelSearch) {
	errst := sv.Validate.Struct(level)
	helper.Err(errst)

	onchan := make(chan bool)
	tx, err := sv.DB.Begin()
	helper.Err(err)

	defer helper.Tx(tx)
	defer close(onchan)

	go sv.LevelRepository.Delete(ctx, tx, domain.Level{
		Id_level: level.Id,
	}, onchan)

	<-onchan
}

// FIndAll implements LevelServices.
func (sv *LevelServicesImpl) FIndAll(ctx context.Context) []response.LevelResponse {
	tx, err := sv.DB.Begin()
	helper.Err(err)
	defer helper.Tx(tx)
	chann := make(chan []domain.Level)
	go sv.LevelRepository.FindAll(ctx, tx, chann)

	res := helper.LevelResponses(<-chann)

	defer close(chann)
	return res

}

// FindById implements LevelServices.
func (sv *LevelServicesImpl) FindById(ctx context.Context, id request.LevelSearch) response.LevelResponse {
	onchann := make(chan domain.Level)
	tx, err := sv.DB.Begin()
	helper.Err(err)

	defer close(onchann)
	defer helper.Tx(tx)

	go sv.LevelRepository.FindById(ctx, tx, id.Id, onchann)
	res := <-onchann
	return helper.LevelResponse(res)
}

// Save implements LevelServices.
func (sv *LevelServicesImpl) Save(ctx context.Context, level request.LevelRequest) {
	errst := sv.Validate.Struct(level)
	helper.Err(errst)

	onchan := make(chan bool)
	tx, err := sv.DB.Begin()
	helper.Err(err)

	go sv.LevelRepository.Save(ctx, tx, domain.Level{Nama_level: level.Nama_level}, onchan)

	defer close(onchan)
	defer helper.Tx(tx)
	<-onchan
}

// Update implements LevelServices.
func (sv *LevelServicesImpl) Update(ctx context.Context, level request.LevelUpdate) {
	errst := sv.Validate.Struct(level)
	helper.Err(errst)

	onchan := make(chan bool)
	tx, err := sv.DB.Begin()
	helper.Err(err)

	go sv.LevelRepository.Update(ctx, tx, domain.Level{Id_level: level.Id_level, Nama_level: level.Name}, onchan)

	defer close(onchan)
	defer helper.Tx(tx)

	<-onchan
}

func NewLevelServices(levelRepository repository.LevelRepository, db *sql.DB, validate *validator.Validate) LevelServices {
	return &LevelServicesImpl{
		LevelRepository: levelRepository,
		DB:              db,
		Validate:        validate,
	}
}
