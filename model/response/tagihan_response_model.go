package response

type TagihanResponse struct {
	Bulan          string `validate:"required" json:"bulan"`
	Tahun          string `validate:"required" json:"tahun"`
	Penggunaan_Kwh int    `validate:"required" json:"penggunaan_kwh"`
	Status         bool   `validate:"required" json:"status"`
}
