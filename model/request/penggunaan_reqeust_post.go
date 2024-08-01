package request

type PenggunaanAdd struct {
	Id_pelanggan int    `validate:"required" json:"id"`
	Bulan        string `validate:"required" json:"bulan"`
	Tahun        string `validate:"required" json:"tahun"`
	Meter_awal   int    `validate:"required" json:"meter_awal"`
	Meter_akhir  int    `validate:"required" json:"meter_akhir"`
}
