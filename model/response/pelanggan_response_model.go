package response

type PelangganResponse struct {
	Id_pelanggan   int32  `json:"id"`
	Username       string `json:"username"`
	Name_pelanggan string `json:"name"`
	Nomor_kwh      string `json:"nomor_kwh"`
	Alamat         string `json:"alamat"`
	Tarif_name     string `json:"tarif"`
}
