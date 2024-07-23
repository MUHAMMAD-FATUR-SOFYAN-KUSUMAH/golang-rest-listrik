package request

type LevelRequest struct {
	Nama_level string `validate:"required" json:"name"`
}
