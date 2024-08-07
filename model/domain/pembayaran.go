package domain

import "time"

type Pembayaran struct {
	Id_pembayaran      int
	Admin_name         string
	No_kwh             string
	Nama_pelanggan     string
	Tanggal_pembayaran time.Time
	Total_bayar        int
	Name_image         string
	Status             int
	Meter_awal         int
	Meter_ahkir        int
	Total_meter        int
	Biaya_admin        int
	Bulan              string
	Tahun              string
}
