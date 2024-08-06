package repository

import (
	"context"
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/domain"
)

type PembayaranRepositoryImpl struct{}

// Delete implements PembayaranRepository.
func (*PembayaranRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pembayaran domain.Pembayaran) {
	panic("unimplemented")
}

// FindAllDetails implements PembayaranRepository.
func (*PembayaranRepositoryImpl) FindAllDetails(ctx context.Context, tx *sql.Tx, done chan []domain.Pembayaran) {

}

// FindAllKonfrimasi implements PembayaranRepository.
func (*PembayaranRepositoryImpl) FindAllKonfrimasi(ctx context.Context, tx *sql.Tx, done chan []domain.Pembayaran) {
	sql := "SELECT pm.id_pembayaran, pm.tanggal_pembayaran, pm.total_bayar, pg.nama_pelanggan ,pg.nomor_kwh, tg.image, tg.status FROM public.pembayaran AS pm JOIN public.tagihan AS tg ON pm.tagihan = tg.id_tagihan JOIN public.pelanggan AS pg ON tg.pelanggan = pg.id_pelanggan WHERE tg.status = 2"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)
	var temp []domain.Pembayaran

	for row.Next() {
		var gory domain.Pembayaran
		err := row.Scan(&gory.Id_pembayaran, &gory.Tanggal_pembayaran, &gory.Total_bayar, &gory.Nama_pelanggan, &gory.No_kwh, &gory.Name_image, &gory.Status)
		helper.Err(err)
		temp = append(temp, gory)
	}

	defer func() {
		done <- temp
	}()
}

func NewPembayaranRepository() PembayaranRepository {
	return &PembayaranRepositoryImpl{}
}
