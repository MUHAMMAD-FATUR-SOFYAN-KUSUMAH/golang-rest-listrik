package domain

type Penggunaan struct {
	Id_pengunaan int
	Id_pelanggan int
	Meter_awal   int
	Meter_ahkir  int
	Bulan        string
	Tahun        string
}
