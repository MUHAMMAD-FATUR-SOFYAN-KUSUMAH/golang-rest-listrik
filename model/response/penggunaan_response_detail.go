package response

type PenggunaanResponseDetail struct {
	IdPenggunaan int    `json:"id_penggunaan"`
	Meter_awal   int    `json:"meter_awal"`
	Meter_akhir  int    `json:"meter_akhir"`
	Bulan        string `json:"bulan"`
	Tahun        string `json:"tahun"`
	Total_Meter  int    `json:"total_meter"`
}
