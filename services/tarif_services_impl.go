package services

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/repository"

	"github.com/go-playground/validator/v10"
)

type TarifServicesImpl struct {
	TarifRepository repository.TarifRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

// Delete implements TarifServices.
func (s *TarifServicesImpl) Delete(ctx context.Context, tarif request.TarifSearch) {
	panic("unimplemented")
}

// FIndAll implements TarifServices.
func (s *TarifServicesImpl) FIndAll(ctx context.Context) []response.TarifResponse {
	tx, err := s.DB.Begin()
	helper.Err(err)
	defer helper.Tx(tx)

	tarif := s.TarifRepository.FindAll(ctx, tx)
	return helper.TarifResponses(tarif)
}

// FindById implements TarifServices.
func (s *TarifServicesImpl) FindById(ctx context.Context, id request.TarifSearch) response.TarifResponse {
	panic("unimplemented")
}

// Save implements TarifServices.
func (s *TarifServicesImpl) Save(ctx context.Context, tarif request.TarifSave) {
	panic("unimplemented")
}

// Update implements TarifServices.
func (s *TarifServicesImpl) Update(ctx context.Context, tarif request.TarifUpdated) {
	panic("unimplemented")
}

func NewTarifServices(Tarif repository.TarifRepository, db *sql.DB, validate *validator.Validate) TarifServices {
	return &TarifServicesImpl{
		TarifRepository: Tarif,
		DB:              db,
		Validate:        validate,
	}
}
