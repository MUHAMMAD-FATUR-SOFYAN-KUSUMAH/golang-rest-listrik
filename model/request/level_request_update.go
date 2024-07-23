package request

type LevelUpdate struct {
	Id_level int    `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
}
