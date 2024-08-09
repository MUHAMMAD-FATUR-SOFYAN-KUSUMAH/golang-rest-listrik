package helper

import (
	"fmt"
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

func PenggunaanIndexPage(penggunaan domain.Pelanggan) response.PenggunaanResponseIndex {
	return response.PenggunaanResponseIndex{
		Id:        penggunaan.Id_pelanggan,
		Nama:      penggunaan.Name_pelanggan,
		Nomor_kwh: penggunaan.Nomor_kwh,
		Alamat:    penggunaan.Alamat,
	}
}

func PenggunaanIndexPages(penggunaan []domain.Pelanggan) []response.PenggunaanResponseIndex {
	var temp []response.PenggunaanResponseIndex
	for _, penggunaan := range penggunaan {
		temp = append(temp, PenggunaanIndexPage(penggunaan))
	}
	return temp
}

func PenggunaanDetail(penggunaan domain.Penggunaan) response.PenggunaanResponseDetail {
	return response.PenggunaanResponseDetail{
		IdPenggunaan: penggunaan.Id_pengunaan,
		Meter_awal:   penggunaan.Meter_awal,
		Meter_akhir:  penggunaan.Meter_ahkir,
		Bulan:        penggunaan.Bulan,
		Tahun:        penggunaan.Tahun,
		Total_Meter:  penggunaan.Meter_ahkir - penggunaan.Meter_awal,
	}
}

func PenggunaanDeatils(penggunaan []domain.Penggunaan) []response.PenggunaanResponseDetail {
	var temp []response.PenggunaanResponseDetail
	for _, penggunaan := range penggunaan {
		temp = append(temp, PenggunaanDetail(penggunaan))
	}
	return temp
}

func PembayaranConf(pembayaran domain.Pembayaran) response.PembayaranConfResponse {
	Res_role := Int_StatusToString(pembayaran.Status)
	bulan_bayar := fmt.Sprintf("%s %s", pembayaran.Bulan, pembayaran.Tahun)
	TotalBayar := fmt.Sprintf("Rp. %d", pembayaran.Total_bayar)

	return response.PembayaranConfResponse{
		Id_pembayaran:  pembayaran.Id_pembayaran,
		No_kwh:         pembayaran.No_kwh,
		Bulan_bayar:    bulan_bayar,
		Name_Pelanggan: pembayaran.Nama_pelanggan,
		Total_Bayar:    TotalBayar,
		Name_image:     pembayaran.Name_image,
		Status:         Res_role,
		Tagihan_prim:   pembayaran.Id_tagihan,
	}
}

func PembayaranConfs(pembayaran []domain.Pembayaran) []response.PembayaranConfResponse {
	var temp []response.PembayaranConfResponse
	for _, pembayaran := range pembayaran {
		temp = append(temp, PembayaranConf(pembayaran))
	}
	return temp
}

func PembayaranDetail(pembayaran domain.Pembayaran) response.PembayaranDetailResponse {
	bulan_bayar := fmt.Sprintf("%s %s", pembayaran.Bulan, pembayaran.Tahun)
	Res_role := Int_StatusToString(pembayaran.Status)
	NewFormat := pembayaran.Tanggal_pembayaran.Format("02-01-2006")
	return response.PembayaranDetailResponse{
		Id_pembayaran:  pembayaran.Id_pembayaran,
		No_kwh:         pembayaran.No_kwh,
		Name_Pelanggan: pembayaran.Nama_pelanggan,
		Tanggal_bayar:  NewFormat,
		Total_Bayar:    pembayaran.Total_bayar,
		Name_image:     pembayaran.Name_image,
		Status:         Res_role,
		Meter_Awal:     pembayaran.Meter_awal,
		Meter_Ahkir:    pembayaran.Meter_ahkir,
		Total_Meter:    pembayaran.Total_meter,
		Admin_Pemverif: pembayaran.Admin_name,
		Bulan_bayar:    bulan_bayar,
	}
}

func PembayaranDetails(pembayaran []domain.Pembayaran) []response.PembayaranDetailResponse {
	var temp []response.PembayaranDetailResponse
	for _, pembayaran := range pembayaran {
		temp = append(temp, PembayaranDetail(pembayaran))
	}
	return temp
}

func LoginResponse(user domain.User) response.LoginResponse {
	status := Generate_role(user.Level)
	return response.LoginResponse{
		Id:   int(user.Id_user),
		Name: user.Name,
		Role: status,
	}
}
