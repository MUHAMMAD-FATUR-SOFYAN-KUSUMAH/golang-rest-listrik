package controller

import (
	"fmt"
	"golang_listrik/helper"
	"golang_listrik/model/request"
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
func (controller *tarifControllerImpl) Deleted(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Temp := request.TarifSearch{}
	helper.Decode_Json(r, &Temp)

	ctx := r.Context()
	controller.Services.Delete(ctx, Temp)

	Response := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeth",
		Data:   "behasil menghapus",
	}

	helper.Encode_Json(w, Response)
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
func (controller *tarifControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	TarifCreate := request.TarifSave{}
	helper.Decode_Json(r, &TarifCreate)

	controller.Services.Save(r.Context(), TarifCreate)

	WebResponse := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accept",
		Data:   "success save",
	}

	helper.Encode_Json(w, WebResponse)
}

// Updated implements TarifController.
func (controller *tarifControllerImpl) Updated(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	chann := make(chan bool)
	temp := request.TarifUpdated{}
	helper.Decode_Json(r, &temp)

	go controller.Services.Update(r.Context(), temp, chann)

	WebResponse := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accept",
		Data:   "success save",
	}

	helper.Encode_Json(w, WebResponse)
	defer func() {
		<-chann
	}()
}

func NewTarifControllerImpl(tarifServices services.TarifServices) TarifController {
	return &tarifControllerImpl{
		Services: tarifServices,
	}
}
