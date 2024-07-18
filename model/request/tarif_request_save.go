package request

type TarifSave struct {
	Daya     int `validate:"required" json:"daya"`
	TarifKwh int `validate:"required" json:"Tarif"`
}
