package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type PelangganRepositoryImpl struct {
}

// Delete implements PelangganRepository.
func (*PelangganRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) {
	panic("unimplemented")
}

// FindAll implements PelangganRepository.
func (*PelangganRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Pelanggan) {
	sql := "SELECT p.id_pelanggan, p.username, p.nama_pelanggan, p.nomor_kwh, p.alamat, p.tarif ,t.nama_tarif ,t.id_tarif FROM public.pelanggan AS p INNER JOIN public.tarif as t ON p.tarif = t.id_tarif"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	defer row.Close()

	var model_pelanggan []domain.Pelanggan
	for row.Next() {
		var gory domain.Pelanggan
		tempp := 0
		err := row.Scan(&gory.Id_pelanggan, &gory.Username, &gory.Name_pelanggan, &gory.Nomor_kwh, &gory.Alamat, &gory.Tarif_id, &gory.Tarif.Name, &tempp)
		helper.Err(err)
		model_pelanggan = append(model_pelanggan, gory)
	}

	defer func() {
		done <- model_pelanggan
	}()

}

// FindById implements PelangganRepository.
func (*PelangganRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Pelanggan) {
	sql := "SELECT p.id_pelanggan, p.username, p.nama_pelanggan, p.nomor_kwh, p.alamat, p.tarif ,t.nama_tarif ,t.id_tarif FROM public.pelanggan AS p INNER JOIN public.tarif as t ON p.tarif = t.id_tarif WHERE p.id_pelanggan = $1"
	row := tx.QueryRowContext(ctx, sql, id)
	var pelanggan domain.Pelanggan
	tempp := 0
	err := row.Scan(&pelanggan.Id_pelanggan, &pelanggan.Username, &pelanggan.Name_pelanggan, &pelanggan.Nomor_kwh, &pelanggan.Alamat, &pelanggan.Tarif_id, &pelanggan.Tarif.Name, &tempp)
	helper.Err(err)
	chann <- pelanggan
}

// Save implements PelangganRepository.
func (*PelangganRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) {
	panic("unimplemented")
}

// Update implements PelangganRepository.
func (*PelangganRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) {
	panic("unimplemented")
}

func NewPelangganRepsitoryImpl() PelangganRepository {
	return &PelangganRepositoryImpl{}
}
