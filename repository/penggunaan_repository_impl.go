package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type PenggunaanRepositoryImpl struct{}

// FindAll implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, done chan []domain.Pelanggan) {
	sql := "SELECT id_pelanggan, nama_pelanggan, alamat, nomor_kwh FROM public.pelanggan"

	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	var model []domain.Pelanggan

	for row.Next() {
		var gory domain.Pelanggan
		err := row.Scan(&gory.Id_pelanggan, &gory.Name_pelanggan, &gory.Alamat, &gory.Nomor_kwh)
		helper.Err(err)
		model = append(model, gory)
	}

	defer func() {
		done <- model
		row.Close()
	}()

}

// Delete implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan) {
	sql_2 := "DELETE FROM public.tagihan WHERE penggunaan = $1"
	_, err_2 := tx.ExecContext(ctx, sql_2, penggunaan.Id_pengunaan)
	helper.Err(err_2)

	sql := "DELETE FROM public.penggunaan WHERE id_penggunaan = $1"
	_, err := tx.ExecContext(ctx, sql, penggunaan.Id_pengunaan)
	helper.Err(err)

}

// FindAll implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) FindAllDetail(ctx context.Context, tx *sql.Tx, done chan []domain.Penggunaan, id int) {
	sql := "SELECT id_penggunaan, bulan, tahun , meter_awal, meter_ahkir FROM public.penggunaan WHERE pelanggan = $1"

	row, err := tx.QueryContext(ctx, sql, id)
	helper.Err(err)

	var model []domain.Penggunaan

	for row.Next() {
		var gory domain.Penggunaan
		err := row.Scan(&gory.Id_pengunaan, &gory.Bulan, &gory.Tahun, &gory.Meter_awal, &gory.Meter_ahkir)
		helper.Err(err)
		model = append(model, gory)
	}

	defer func() {
		done <- model
		row.Close()
	}()
}

// FindById implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, chann chan domain.Penggunaan) {
	sql := "SELECT id_penggunaan, bulan, tahun , meter_awal, meter_ahkir FROM public.penggunaan WHERE id_penggunaan = $1"
	row := tx.QueryRowContext(ctx, sql, id)
	temp := domain.Penggunaan{}
	err := row.Scan(&temp.Id_pengunaan, &temp.Bulan, &temp.Tahun, &temp.Meter_awal, &temp.Meter_ahkir)
	helper.Err(err)

	defer func() {
		chann <- temp
	}()
}

// Save implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan) {
	var id int
	sql := "INSERT INTO public.penggunaan (pelanggan, bulan, tahun, meter_awal, meter_ahkir) VALUES ($1, $2, $3, $4, $5) RETURNING id_penggunaan"
	row := tx.QueryRowContext(ctx, sql, penggunaan.Id_pelanggan, penggunaan.Bulan, penggunaan.Tahun, penggunaan.Meter_awal, penggunaan.Meter_ahkir).Scan(&id)
	helper.Err(row)

	total := penggunaan.Meter_ahkir - penggunaan.Meter_awal

	// insert tagihan
	sql2 := "INSERT INTO public.tagihan (penggunaan, pelanggan, jumlah_meter) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, sql2, id, penggunaan.Id_pelanggan, total)

	helper.Err(err)
}

// Update implements PenggunaanRepository.
func (*PenggunaanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, penggunaan domain.Penggunaan) {
	sql := "UPDATE public.penggunaan SET bulan = $1, tahun = $2, meter_awal = $3, meter_ahkir = $4 WHERE id_penggunaan = $5"
	_, err := tx.ExecContext(ctx, sql, penggunaan.Bulan, penggunaan.Tahun, penggunaan.Meter_awal, penggunaan.Meter_ahkir, penggunaan.Id_pengunaan)
	helper.Err(err)

	reuslt := penggunaan.Meter_ahkir - penggunaan.Meter_awal

	// insert tagihan
	sql2 := "UPDATE public.tagihan SET jumlah_meter = $1 WHERE penggunaan = $2"
	_, err_updated := tx.Exec(sql2, reuslt, penggunaan.Id_pengunaan)
	helper.Err(err_updated)
}

func NewPenggunaanRepositoryImpl() PenggunaanRepository {
	return &PenggunaanRepositoryImpl{}
}
