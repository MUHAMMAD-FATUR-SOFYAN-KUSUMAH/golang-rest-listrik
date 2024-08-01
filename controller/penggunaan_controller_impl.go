package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PenggunaanControllerImpl struct {
	penggunaanService services.PenggunaanServices
}

// Delete implements PenggunaanController.
func (controller *PenggunaanControllerImpl) Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.PenggunaanSearch{}
	ch := make(chan bool)
	helper.Decode_Json(r, &temp)

	go controller.penggunaanService.Delete(r.Context(), temp, ch)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Data deleted",
	}

	<-ch

	helper.Encode_Json(writer, webResponse)
}

// FindAll implements PenggunaanController.
func (controller *PenggunaanControllerImpl) FindAll(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := controller.penggunaanService.FindAll(r.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.Encode_Json(writer, webResponse)
}

// FindAllDetail implements PenggunaanController.
func (controller *PenggunaanControllerImpl) FindAllDetail(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id request.PelangganRequest

	helper.Decode_Json(r, &id)
	result := controller.penggunaanService.FindAllDetail(r.Context(), id.Id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.Encode_Json(writer, webResponse)
}

// FindById implements PenggunaanController.
func (controller *PenggunaanControllerImpl) FindById(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.PenggunaanSearch{}
	helper.Decode_Json(r, &temp)
	chann := make(chan response.PenggunaanResponseDetail)

	go controller.penggunaanService.FindById(r.Context(), temp.Id, chann)

	model := <-chann

	custome := response.PenggunaanResponseDetail{
		IdPenggunaan: model.IdPenggunaan,
		Meter_awal:   model.Meter_awal,
		Meter_akhir:  model.Meter_akhir,
		Bulan:        model.Bulan,
		Tahun:        model.Tahun,
	}
	close(chann)
	helper.Encode_Json(writer, custome)
}

// Save implements PenggunaanController.
func (controller *PenggunaanControllerImpl) Save(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var penggunaan request.PenggunaanAdd
	helper.Decode_Json(r, &penggunaan)

	controller.penggunaanService.Save(r.Context(), penggunaan)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "data saved",
	}

	helper.Encode_Json(writer, webResponse)
}

// Update implements PenggunaanController.
func (controller *PenggunaanControllerImpl) Update(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var temp request.PenggunaanUpdate

	helper.Decode_Json(r, &temp)

	controller.penggunaanService.Update(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "data updated",
	}

	helper.Encode_Json(writer, webResponse)
}

func NewPenggunaanController(penggunaanService services.PenggunaanServices) PenggunaanController {
	return &PenggunaanControllerImpl{
		penggunaanService: penggunaanService,
	}
}
