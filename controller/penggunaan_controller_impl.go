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
func (*PenggunaanControllerImpl) Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
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
func (*PenggunaanControllerImpl) FindById(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Save implements PenggunaanController.
func (*PenggunaanControllerImpl) Save(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Update implements PenggunaanController.
func (*PenggunaanControllerImpl) Update(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

func NewPenggunaanController(penggunaanService services.PenggunaanServices) PenggunaanController {
	return &PenggunaanControllerImpl{
		penggunaanService: penggunaanService,
	}
}
