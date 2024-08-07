package response

type PembayaranDetailResponse struct {
	Id_pembayaran  int    `json:"id_tagihan"`
	No_kwh         string `validate:"required" json:"no_kwh"`
	Name_Pelanggan string `validate:"required" json:"name_pelanggan"`
	Tanggal_bayar  string `validate:"required" json:"tanggal_bayar"`
	Total_Bayar    any    `validate:"required" json:"total_bayar"`
	Name_image     string `validate:"required" json:"name_image"`
	Status         any    `validate:"required" json:"status"`
	Bulan_bayar    string `validate:"required" json:"bulan_bayar"`
	Meter_Awal     int    `validate:"required" json:"meter_awal"`
	Meter_Ahkir    int    `validate:"required" json:"meter_ahkir"`
	Total_Meter    int    `validate:"required" json:"total_meter"`
	Admin_Pemverif string `validate:"required" json:"admin_pemverif"`
}
