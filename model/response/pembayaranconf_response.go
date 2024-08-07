package response

type PembayaranConfResponse struct {
	Id_pembayaran  int    `validate:"required" json:"id_tagihan"`
	No_kwh         string `validate:"required" json:"no_kwh"`
	Name_Pelanggan string `validate:"required" json:"name_pelanggan"`
	Bulan_bayar    string `validate:"required" json:"bulan_bayar"`
	Total_Bayar    any    `validate:"required" json:"total_bayar"`
	Name_image     string `validate:"required" json:"name_image"`
	Status         any    `validate:"required" json:"status"`
}
