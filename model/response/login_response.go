package response

type LoginResponse struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Id   int    `json:"id"`
}
