package response

type PenggunaanResponseIndex struct {
	Id        int    `validate:"required" json:"id"`
	Nama      string `validate:"required" json:"nama"`
	Nomor_kwh string `validate:"required" json:"nomor_kwh"`
	Alamat    string `validate:"required" json:"alamat"`
}
