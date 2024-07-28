package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PelangganControllerImpl struct {
	pelangganservices services.PelangganServices
}

// Delete implements PelangganController.
func (*PelangganControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// FindAll implements PelangganController.
func (controller *PelangganControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	result := controller.pelangganservices.FindAll(request.Context())

	webrespons := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   result,
	}

	helper.Encode_Json(writer, webrespons)
}

// FindById implements PelangganController.
func (controller *PelangganControllerImpl) FindById(writer http.ResponseWriter, req *http.Request, p httprouter.Params) {
	temp := request.PelangganRequest{}
	helper.Decode_Json(req, &temp)

	result := controller.pelangganservices.FindById(req.Context(), temp.Id)

	webrespons := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   result,
	}

	helper.Encode_Json(writer, webrespons)
}

// Save implements PelangganController.
func (*PelangganControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// Update implements PelangganController.
func (*PelangganControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

func NewPelangganControllerImpl(pelangganservices services.PelangganServices) PelangganController {
	return &PelangganControllerImpl{
		pelangganservices: pelangganservices,
	}
}
