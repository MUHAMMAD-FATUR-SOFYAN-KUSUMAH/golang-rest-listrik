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
	sql := "SELECT pb.tanggal_pembayaran, pb.biaya_admin, pb.total_bayar, tg.jumlah_meter , tg.status , tg.image , pn.meter_awal, pn.meter_ahkir, pn.bulan, pn.tahun, us.nama_admin, pg.nama_pelanggan FROM public.pembayaran AS pb JOIN public.user AS us ON pb.user = us.id_user JOIN public.tagihan AS tg ON pb.tagihan = tg.id_tagihan JOIN public.penggunaan AS pn ON tg.penggunaan = pn.id_penggunaan JOIN public.pelanggan AS pg ON pn.pelanggan = pg.id_pelanggan"

	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)

	var temp []domain.Pembayaran

	for row.Next() {
		var gory domain.Pembayaran
		row.Scan(&gory.Tanggal_pembayaran, &gory.Biaya_admin, &gory.Total_bayar, &gory.Total_meter, &gory.Status, &gory.Name_image, &gory.Meter_awal, &gory.Meter_ahkir, &gory.Bulan, &gory.Tahun, &gory.Admin_name, &gory.Nama_pelanggan)
		temp = append(temp, gory)
	}

	defer func() {
		done <- temp
		row.Close()
	}()
}

// FindAllKonfrimasi implements PembayaranRepository.
func (*PembayaranRepositoryImpl) FindAllKonfrimasi(ctx context.Context, tx *sql.Tx, done chan []domain.Pembayaran) {
	sql := "SELECT pm.id_pembayaran, pm.total_bayar, pg.nama_pelanggan ,pg.nomor_kwh, tg.image, tg.status , pn.bulan, pn.tahun FROM public.pembayaran AS pm JOIN public.tagihan AS tg ON pm.tagihan = tg.id_tagihan JOIN public.pelanggan AS pg ON tg.pelanggan = pg.id_pelanggan JOIN public.penggunaan AS pn ON tg.penggunaan = pn.id_penggunaan WHERE tg.status = 2"
	row, err := tx.QueryContext(ctx, sql)
	helper.Err(err)
	var temp []domain.Pembayaran

	for row.Next() {
		var gory domain.Pembayaran
		err := row.Scan(&gory.Id_pembayaran, &gory.Total_bayar, &gory.Nama_pelanggan, &gory.No_kwh, &gory.Name_image, &gory.Status, &gory.Bulan, &gory.Tahun)
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
