package request

type UpdatePelanggan struct {
	Id_pelanggan int    `validate:"required" json:"id"`
	Username     string `validate:"required" json:"username"`
	Password     string `validate:"required" json:"password"`
	Name         string `validate:"required" json:"name"`
	Alamat       string `validate:"required" json:"alamat"`
	Tarif_id     int    `validate:"required" json:"tarif_id"`
}
