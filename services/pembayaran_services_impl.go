package services

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/repository"
)

type PembayaranServicesImpl struct {
	DB           *sql.DB
	RepositoryPb repository.PembayaranRepository
}

// Delete implements PembayaranServices.
func (controller *PembayaranServicesImpl) Delete(ctx context.Context, pembayaran request.LevelSearch) {
	tx, _ := controller.DB.Begin()

	controller.RepositoryPb.Delete(ctx, tx, domain.Pembayaran{
		Id_tagihan: pembayaran.Id,
	})
	helper.Tx(tx)
}

// FindAllDetails implements PembayaranServices.
func (services *PembayaranServicesImpl) FindAllDetails(ctx context.Context) []response.PembayaranDetailResponse {
	tx, _ := services.DB.Begin()
	chann := make(chan []domain.Pembayaran)
	go services.RepositoryPb.FindAllDetails(ctx, tx, chann)

	resonse_model := helper.PembayaranDetails(<-chann)
	defer func() {
		helper.Tx(tx)
		close(chann)
	}()

	return resonse_model
}

// FindAllKonfrimasi implements PembayaranServices.
func (services *PembayaranServicesImpl) FindAllKonfrimasi(ctx context.Context) []response.PembayaranConfResponse {
	tx, _ := services.DB.Begin()
	Onichan := make(chan []domain.Pembayaran)
	go services.RepositoryPb.FindAllKonfrimasi(ctx, tx, Onichan)

	resonse_model := helper.PembayaranConfs(<-Onichan)

	defer func() {
		helper.Tx(tx)
		close(Onichan)
	}()

	return resonse_model
}

func NewPembayaranServices8(db *sql.DB, repositoryPb repository.PembayaranRepository) PembayaranServices {
	return &PembayaranServicesImpl{
		DB:           db,
		RepositoryPb: repositoryPb,
	}
}
