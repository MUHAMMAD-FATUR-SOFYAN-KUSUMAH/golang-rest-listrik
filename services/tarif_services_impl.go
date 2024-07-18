package services

import (
	"context"
	"database/sql"
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
func (*TarifServicesImpl) Delete(ctx context.Context, tarif request.TarifSearch) {
	panic("unimplemented")
}

// FIndAll implements TarifServices.
func (*TarifServicesImpl) FIndAll(ctx context.Context) []response.TarifResponse {
	panic("unimplemented")
}

// FindById implements TarifServices.
func (*TarifServicesImpl) FindById(ctx context.Context, id request.TarifSearch) response.TarifResponse {
	panic("unimplemented")
}

// Save implements TarifServices.
func (*TarifServicesImpl) Save(ctx context.Context, tarif request.TarifSave) {
	panic("unimplemented")
}

// Update implements TarifServices.
func (*TarifServicesImpl) Update(ctx context.Context, tarif request.TarifUpdated) {
	panic("unimplemented")
}

func NewTarifServices(Tarif repository.TarifRepository, db *sql.DB, validate *validator.Validate) TarifServices {
	return &TarifServicesImpl{
		TarifRepository: Tarif,
		DB:              db,
		Validate:        validate,
	}
}
