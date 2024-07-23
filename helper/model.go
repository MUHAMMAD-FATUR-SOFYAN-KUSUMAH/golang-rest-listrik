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
