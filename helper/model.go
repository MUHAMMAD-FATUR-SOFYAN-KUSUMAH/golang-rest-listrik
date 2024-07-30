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

func PelangganReponse(pelanggan domain.Pelanggan) response.PelangganResponse {
	return response.PelangganResponse{
		Id_pelanggan:   int32(pelanggan.Id_pelanggan),
		Username:       pelanggan.Username,
		Name_pelanggan: pelanggan.Name_pelanggan,
		Nomor_kwh:      pelanggan.Nomor_kwh,
		Alamat:         pelanggan.Alamat,
		Tarif_name:     pelanggan.Tarif.Name,
	}
}

func PelangganResponses(pelanggan []domain.Pelanggan) []response.PelangganResponse {
	var temp []response.PelangganResponse
	for _, pelanggan := range pelanggan {
		temp = append(temp, PelangganReponse(pelanggan))
	}
	return temp
}

func PenggunaanResponse(penggunaan domain.Penggunaan) response.PenggunaanResponseDetail {
	total_kwh := penggunaan.Meter_awal - penggunaan.Meter_ahkir
	return response.PenggunaanResponseDetail{
		IdPenggunaan: penggunaan.Id_pengunaan,
		Meter_awal:   penggunaan.Meter_awal,
		Meter_akhir:  penggunaan.Meter_ahkir,
		Bulan:        penggunaan.Bulan,
		Tahun:        penggunaan.Tahun,
		Total_Meter:  total_kwh,
	}
}

func PenggunaanResponses(penggunaan []domain.Penggunaan) []response.PenggunaanResponseDetail {
	var temp []response.PenggunaanResponseDetail
	for _, penggunaan := range penggunaan {
		temp = append(temp, PenggunaanResponse(penggunaan))
	}
	return temp
}
