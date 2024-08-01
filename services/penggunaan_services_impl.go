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
func (services *PenggunaanServicesImpl) FindById(ctx context.Context, id int, chann chan response.PenggunaanResponseDetail) {
	maken := make(chan domain.Penggunaan)
	tx, _ := services.Db.Begin()

	go services.PenggunaanRepository.FindById(ctx, tx, id, maken)
	res := helper.PenggunaanDetail(<-maken)

	chann <- res
	defer func() {
		helper.Tx(tx)
		close(maken)
	}()
}

// Save implements PenggunaanServices.
func (services *PenggunaanServicesImpl) Save(ctx context.Context, penggunaan request.PenggunaanAdd) {
	err_validator := services.validat.Struct(penggunaan)
	tx, _ := services.Db.Begin()
	helper.Err(err_validator)

	temp := domain.Penggunaan{
		Id_pelanggan: penggunaan.Id_pelanggan,
		Bulan:        penggunaan.Bulan,
		Tahun:        penggunaan.Tahun,
		Meter_awal:   penggunaan.Meter_awal,
		Meter_ahkir:  penggunaan.Meter_akhir,
	}

	services.PenggunaanRepository.Save(ctx, tx, temp)
}

// Update implements PenggunaanServices.
func (services *PenggunaanServicesImpl) Update(ctx context.Context, penggunaan request.PenggunaanUpdate) {
	err_validator_struct := services.validat.Struct(penggunaan)
	tx, _ := services.Db.Begin()
	helper.Err(err_validator_struct)

	services.PenggunaanRepository.Update(ctx, tx, domain.Penggunaan{
		Id_pengunaan: penggunaan.Id,
		Meter_awal:   penggunaan.Meter_awal,
		Meter_ahkir:  penggunaan.Meter_akhir,
		Bulan:        penggunaan.Bulan,
		Tahun:        penggunaan.Tahun,
	})
}

func NewPenggunaanServices(validator *validator.Validate, db *sql.DB, Penggunaan repository.PenggunaanRepository) PenggunaanServices {
	return &PenggunaanServicesImpl{
		validat:              validator,
		Db:                   db,
		PenggunaanRepository: Penggunaan,
	}
}
