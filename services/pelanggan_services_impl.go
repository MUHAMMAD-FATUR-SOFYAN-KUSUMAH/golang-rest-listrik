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

type PelangganServicesImpl struct {
	PelangganRepository repository.PelangganRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

// Delete implements PelangganServices.
func (services *PelangganServicesImpl) Delete(ctx context.Context, pelanggan request.PelangganRequest) {
	err := services.Validate.Struct(pelanggan)
	helper.Err(err)
	tx, errdb := services.DB.Begin()
	helper.Err(errdb)
	defer helper.Tx(tx)
	services.PelangganRepository.Delete(ctx, tx, domain.Pelanggan{Id_pelanggan: pelanggan.Id})
}

// FindAll implements PelangganServices.
func (s *PelangganServicesImpl) FindAll(ctx context.Context) []response.PelangganResponse {
	chann := make(chan []domain.Pelanggan)
	tx, err := s.DB.Begin()
	helper.Err(err)
	defer func() {
		helper.Tx(tx)
		close(chann)
	}()
	go s.PelangganRepository.FindAll(ctx, tx, chann)
	pelanggan := <-chann
	return helper.PelangganResponses(pelanggan)
}

// FindById implements PelangganServices.
func (s *PelangganServicesImpl) FindById(ctx context.Context, id int) response.PelangganResponse {
	chann := make(chan domain.Pelanggan)
	tx, err := s.DB.Begin()
	helper.Err(err)
	defer func() {
		helper.Tx(tx)
		close(chann)
	}()
	go s.PelangganRepository.FindById(ctx, tx, id, chann)
	pelanggan := <-chann
	return helper.PelangganReponse(pelanggan)
}

// Save implements PelangganServices.
func (services *PelangganServicesImpl) Save(ctx context.Context, pelanggan request.AddPelanggan) {
	tx, _ := services.DB.Begin()
	err := services.Validate.Struct(pelanggan)
	helper.Err(err)

	defer helper.Tx(tx)
	services.PelangganRepository.Save(ctx, tx, domain.Pelanggan{
		Username:       pelanggan.Username,
		Password:       pelanggan.Password,
		Name_pelanggan: pelanggan.Name,
		Alamat:         pelanggan.Alamat,
		Tarif_id:       pelanggan.Tarif_id,
	})
}

// Update implements PelangganServices.
func (services *PelangganServicesImpl) Update(ctx context.Context, pelanggan request.UpdatePelanggan) {
	err := services.Validate.StructCtx(ctx, pelanggan)
	helper.Err(err)

	tx, _ := services.DB.Begin()
	services.PelangganRepository.Update(ctx, tx, domain.Pelanggan{
		Id_pelanggan:   pelanggan.Id_pelanggan,
		Username:       pelanggan.Username,
		Password:       pelanggan.Password,
		Name_pelanggan: pelanggan.Name,
		Alamat:         pelanggan.Alamat,
		Tarif_id:       pelanggan.Tarif_id,
	})
	defer helper.Tx(tx)
}

func NewPelangganServicesImpl(repository repository.PelangganRepository, db *sql.DB, validate *validator.Validate) PelangganServices {
	return &PelangganServicesImpl{
		PelangganRepository: repository,
		DB:                  db,
		Validate:            validate,
	}
}
