package response

type PenggunaanResponseTagihan struct {
	Bulan          string `validate:"required" json:"bulan"`
	Tahun          string `validate:"required" json:"tahun"`
	Penggunaan_Kwh int    `validate:"required" json:"penggunaan_kwh"`
}
