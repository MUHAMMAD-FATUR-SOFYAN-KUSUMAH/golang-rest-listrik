package response

type UserResponse struct {
	Id_user    int32  `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Role_level string `json:"role"`
}
