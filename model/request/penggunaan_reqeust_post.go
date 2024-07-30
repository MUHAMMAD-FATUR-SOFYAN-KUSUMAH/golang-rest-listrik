package request

type PenggunaanAdd struct {
	Nama_pelanggan string `validate:"required" json:"pelanggan"`
	Bulan          string `validate:"required" json:"bulan"`
	Tahun          string `validate:"required" json:"tahun"`
	Meter_awal     int    `validate:"required" json:"meter_awal"`
	Meter_akhir    int    `validate:"required" json:"meter_akhir"`
}
