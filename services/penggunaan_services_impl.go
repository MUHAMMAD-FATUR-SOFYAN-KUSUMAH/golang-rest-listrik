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

type PenggunaanServicesImpl struct {
	validat              *validator.Validate
	Db                   *sql.DB
	PenggunaanRepository repository.PenggunaanRepository
}

// Delete implements PenggunaanServices.
func (*PenggunaanServicesImpl) Delete(ctx context.Context, penggunaan request.PenggunaanSearch) {
	panic("unimplemented")
}

// FindAll implements PenggunaanServices.
func (services *PenggunaanServicesImpl) FindAll(ctx context.Context) []response.PenggunaanResponseIndex {
	tx, _ := services.Db.Begin()
	chann := make(chan []domain.Pelanggan)
	go services.PenggunaanRepository.FindAll(ctx, tx, chann)

	model := <-chann
	defer func() {
		helper.Tx(tx)
		close(chann)
	}()
	return helper.PenggunaanIndexPages(model)
}

// FindAllDetail implements PenggunaanServices.
func (services *PenggunaanServicesImpl) FindAllDetail(ctx context.Context, id int) []response.PenggunaanResponseDetail {
	tx, err := services.Db.Begin()
	chann := make(chan []domain.Penggunaan)
	helper.Err(err)

	go services.PenggunaanRepository.FindAllDetail(ctx, tx, chann, id)

	model := <-chann
	defer func() {
		helper.Tx(tx)
		close(chann)
	}()
	return helper.PenggunaanDeatils(model)
}

// FindById implements PenggunaanServices.
func (*PenggunaanServicesImpl) FindById(ctx context.Context, id int, chann chan response.PenggunaanResponseDetail) {
	panic("unimplemented")
}

// Save implements PenggunaanServices.
func (*PenggunaanServicesImpl) Save(ctx context.Context, penggunaan request.PenggunaanAdd) {
	panic("unimplemented")
}

// Update implements PenggunaanServices.
func (*PenggunaanServicesImpl) Update(ctx context.Context, penggunaan request.PenggunaanUpdate) {
	panic("unimplemented")
}

func NewPenggunaanServices(validator *validator.Validate, db *sql.DB, Penggunaan repository.PenggunaanRepository) PenggunaanServices {
	return &PenggunaanServicesImpl{
		validat:              validator,
		Db:                   db,
		PenggunaanRepository: Penggunaan,
	}
}
