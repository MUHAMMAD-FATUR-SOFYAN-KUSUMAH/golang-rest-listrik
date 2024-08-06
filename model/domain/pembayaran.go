package domain

type Pembayaran struct {
	Id_pembayaran      int
	No_kwh             string
	Nama_pelanggan     string
	Tanggal_pembayaran string
	Total_bayar        int
	Name_image         string
	Status             int
	Meter_awal         int
	Meter_ahkir        int
	Biaya_meter        int
}
