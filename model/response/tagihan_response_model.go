package response

type TagihanResponse struct {
	Id_tagihan     int
	Bulan          string `validate:"required" json:"bulan"`
	Tahun          string `validate:"required" json:"tahun"`
	Penggunaan_Kwh int    `validate:"required" json:"penggunaan_kwh"`
	Status         any    `validate:"required" json:"status"`
}
