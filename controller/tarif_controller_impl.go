package controller

import (
	"fmt"
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"
	"time"

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
	start := time.Now()
	categoryResponses := controller.Services.FIndAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.Encode_Json(writer, webResponse)
	duration := time.Since(start)
	fmt.Printf("exampleFunction took %v\n", duration)
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
