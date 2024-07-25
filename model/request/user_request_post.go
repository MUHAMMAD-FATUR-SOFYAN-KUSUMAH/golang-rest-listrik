package request

type AddUser struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
	Role     string `validate:"required" json:"role"`
}
