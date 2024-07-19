package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type tarifControllerImpl struct {
	Services services.TarifServices
}

// Deleted implements TarifController.
func (*tarifControllerImpl) Deleted(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements TarifController.
func (controller *tarifControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.Services.FIndAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.Encode_Json(writer, webResponse)
}

// FindById implements TarifController.
func (*tarifControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Save implements TarifController.
func (*tarifControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Updated implements TarifController.
func (*tarifControllerImpl) Updated(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

func NewTarifControllerImpl(tarifServices services.TarifServices) TarifController {
	return &tarifControllerImpl{
		Services: tarifServices,
	}
}
