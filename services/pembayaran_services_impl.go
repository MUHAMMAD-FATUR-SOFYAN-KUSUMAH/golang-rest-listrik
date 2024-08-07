package services

import (
	"context"
	"database/sql"
	"golang_listrik/model/domain"
	"golang_listrik/model/request"
	"golang_listrik/repository"
)

type PembayaranServicesImpl struct {
	DB           *sql.DB
	RepositoryPb repository.PembayaranRepository
}

// Delete implements PembayaranServices.
func (*PembayaranServicesImpl) Delete(ctx context.Context, pembayaran request.LevelSearch) {
	panic("unimplemented")
}

// FindAllDetails implements PembayaranServices.
func (services *PembayaranServicesImpl) FindAllDetails(ctx context.Context) {
	tx, _ := services.DB.Begin()
	chann := make(chan []domain.Pembayaran)
	go services.RepositoryPb.FindAllDetails(ctx, tx, chann)

	model := <-chann
}

// FindAllKonfrimasi implements PembayaranServices.
func (*PembayaranServicesImpl) FindAllKonfrimasi(ctx context.Context) {
	panic("unimplemented")
}

func NewPembayaranServices8(db *sql.DB, repositoryPb repository.PembayaranRepository) PembayaranServices {
	return &PembayaranServicesImpl{
		DB:           db,
		RepositoryPb: repositoryPb,
	}
}
