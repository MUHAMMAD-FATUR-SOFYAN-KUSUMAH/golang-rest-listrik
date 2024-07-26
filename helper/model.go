package helper

import (
	"golang_listrik/model/domain"
	"golang_listrik/model/response"
)

func TarifReponse(tarif domain.Tarif) response.TarifResponse {
	return response.TarifResponse{
		Daya:        tarif.Daya,
		TarifPerKwh: tarif.TarifPerKwh,
		UuidTarif:   tarif.UuidTarif,
		Name:        tarif.Name,
	}
}

func TarifResponses(tarif []domain.Tarif) []response.TarifResponse {
	var temp []response.TarifResponse
	for _, Tarif := range tarif {
		temp = append(temp, TarifReponse(Tarif))
	}
	return temp
}

func LevelResponse(lv domain.Level) response.LevelResponse {
	return response.LevelResponse{
		Id_level: lv.Id_level,
		Name:     lv.Nama_level,
	}
}

func LevelResponses(lv []domain.Level) []response.LevelResponse {
	var temp []response.LevelResponse

	for _, result := range lv {
		temp = append(temp, LevelResponse(result))
	}

	return temp
}

func UserReponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		Id_user:    user.Id_user,
		Username:   user.Username,
		Name:       user.Name,
		Role_level: user.Role_level.Nama_level,
	}
}

func UserResponses(user []domain.User) []response.UserResponse {
	var temp []response.UserResponse
	for _, user := range user {
		temp = append(temp, UserReponse(user))
	}
	return temp
}
