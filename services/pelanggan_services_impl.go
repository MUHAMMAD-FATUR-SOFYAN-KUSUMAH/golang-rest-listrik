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
func (*PelangganServicesImpl) Delete(ctx context.Context, pelanggan request.UserSearch) {
	panic("unimplemented")
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
func (*PelangganServicesImpl) Save(ctx context.Context, pelanggan request.AddPelanggan) {
	panic("unimplemented")
}

// Update implements PelangganServices.
func (*PelangganServicesImpl) Update(ctx context.Context, pelanggan request.UpdatePelanggan) {
	panic("unimplemented")
}

func NewPelangganServicesImpl(repository repository.PelangganRepository, db *sql.DB, validate *validator.Validate) PelangganServices {
	return &PelangganServicesImpl{
		PelangganRepository: repository,
		DB:                  db,
		Validate:            validate,
	}
}
