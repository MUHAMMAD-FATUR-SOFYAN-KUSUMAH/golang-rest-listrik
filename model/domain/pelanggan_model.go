package domain

type Pelanggan struct {
	Id_pelanggan   int
	Username       string
	Password       string
	Name_pelanggan string
	Nomor_kwh      string
	Alamat         string
	Tarif_id       int
	Tarif          Tarif
}
