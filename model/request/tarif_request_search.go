package request

type TarifSearch struct {
	Uuid string `validate:"required" json:"id"`
}
