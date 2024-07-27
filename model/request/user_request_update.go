package request

type UserUpdate struct {
	Id_user  int    `validate:"required" json:"id"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
	Role     int    `validate:"required" json:"role"`
}
