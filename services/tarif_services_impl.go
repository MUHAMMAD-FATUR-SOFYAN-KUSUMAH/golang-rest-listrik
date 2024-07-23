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

type TarifServicesImpl struct {
	TarifRepository repository.TarifRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

// Delete implements TarifServices.
func (s *TarifServicesImpl) Delete(ctx context.Context, tarif request.TarifSearch) {
	tx, errdb := s.DB.Begin()
	helper.Err(errdb)

	err := s.Validate.Struct(tarif)
	helper.Err(err)

	search := domain.Tarif{
		UuidTarif: tarif.Uuid,
	}

	defer helper.Tx(tx)

	s.TarifRepository.Delete(ctx, tx, search)
}

// FIndAll implements TarifServices.
func (s *TarifServicesImpl) FIndAll(ctx context.Context) []response.TarifResponse {
	chann := make(chan []domain.Tarif)
	tx, err := s.DB.Begin()
	helper.Err(err)
	defer helper.Tx(tx)
	defer close(chann)
	go s.TarifRepository.FindAll(ctx, tx, chann)

	tarif := <-chann
	return helper.TarifResponses(tarif)
}

// FindById implements TarifServices.
func (s *TarifServicesImpl) FindById(ctx context.Context, id request.TarifSearch) response.TarifResponse {
	tx, _ := s.DB.Begin()
	chann := make(chan domain.Tarif)
	go s.TarifRepository.FindById(ctx, tx, id.Uuid, chann)

	res := helper.TarifReponse(<-chann)

	defer helper.Tx(tx)
	defer close(chann)

	return res
}

// Save implements TarifServices.
func (s *TarifServicesImpl) Save(ctx context.Context, tarif request.TarifSave) {
	err := s.Validate.Struct(tarif)
	helper.Err(err, "lengkapi field")

	tx, ErrDb := s.DB.Begin()
	helper.Err(ErrDb)
	defer helper.Tx(tx)

	ModelTarif := domain.Tarif{
		Daya:        int16(tarif.Daya),
		TarifPerKwh: int32(tarif.TarifKwh),
	}

	s.TarifRepository.Save(ctx, tx, ModelTarif)
}

// Update implements TarifServices.
func (s *TarifServicesImpl) Update(ctx context.Context, tarif request.TarifUpdated, done chan bool) {
	err := s.Validate.StructCtx(ctx, tarif)
	helper.Err(err)

	db, errdb := s.DB.Begin()
	helper.Err(errdb)

	defer helper.Tx(db)

	defer func() {
		done <- true
	}()

	ModelTarif := domain.Tarif{
		Daya:        int16(tarif.Daya),
		TarifPerKwh: int32(tarif.TarifKwh),
		UuidTarif:   tarif.Uuid,
	}

	s.TarifRepository.Update(ctx, db, ModelTarif)

}

func NewTarifServices(Tarif repository.TarifRepository, db *sql.DB, validate *validator.Validate) TarifServices {
	return &TarifServicesImpl{
		TarifRepository: Tarif,
		DB:              db,
		Validate:        validate,
	}
}
