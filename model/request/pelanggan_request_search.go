package request

type PelangganRequest struct {
	Id int `validate:"required" json:"id"`
}
