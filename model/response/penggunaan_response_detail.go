package response

type PenggunaanResponseDetail struct {
	IdPenggunaan int    `validate:"required" json:"id_penggunaan"`
	Meter_awal   int    `validate:"required" json:"meter_awal"`
	Meter_akhir  int    `validate:"required" json:"meter_akhir"`
	Bulan        string `validate:"required" json:"bulan"`
	Tahun        string `validate:"required" json:"tahun"`
	Total_Meter  int    `validate:"required" json:"total_meter"`
}
