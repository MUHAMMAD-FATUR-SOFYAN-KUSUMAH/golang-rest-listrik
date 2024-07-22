package request

type TarifUpdated struct {
	Uuid     string `validate:"required" json:"id"`
	Daya     int    `validate:"required" json:"daya"`
	TarifKwh int    `validate:"required" json:"Tarif"`
}
