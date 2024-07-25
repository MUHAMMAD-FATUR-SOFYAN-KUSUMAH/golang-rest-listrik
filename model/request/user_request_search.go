package request

type UserSearch struct {
	Id int `validate:"required" json:"id"`
}
