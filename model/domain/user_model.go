package domain

type User struct {
	Id_user    int32
	Username   string
	Password   string
	Name       string
	Level      int
	Role_level Level
}
