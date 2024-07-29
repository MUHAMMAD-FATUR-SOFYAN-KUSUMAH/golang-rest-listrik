package request

type AddPelanggan struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
	Alamat   string `validate:"required" json:"alamat"`
	Tarif_id int    `validate:"required" json:"tarif_id"`
}
