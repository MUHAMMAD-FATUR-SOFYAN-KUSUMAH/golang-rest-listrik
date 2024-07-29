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
func (controller *PelangganControllerImpl) Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()

	temp := request.PelangganRequest{}
	helper.Decode_Json(r, &temp)

	controller.pelangganservices.Delete(ctx, temp)

	webrespons := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeth",
		Data:   "data deleted",
	}

	helper.Encode_Json(writer, webrespons)
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
func (controller *PelangganControllerImpl) Save(writer http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := req.Context()
	temp := request.AddPelanggan{}
	helper.Decode_Json(req, &temp)

	controller.pelangganservices.Save(ctx, temp)

	webrespons := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "created",
		Data:   "data Created",
	}
	helper.Encode_Json(writer, webrespons)
}

// Update implements PelangganController.
func (controller *PelangganControllerImpl) Update(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()
	temp := request.UpdatePelanggan{}
	helper.Decode_Json(r, &temp)

	controller.pelangganservices.Update(ctx, temp)

	webrespons := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accepted",
		Data:   "data updated",
	}
	helper.Encode_Json(writer, webrespons)
}

func NewPelangganControllerImpl(pelangganservices services.PelangganServices) PelangganController {
	return &PelangganControllerImpl{
		pelangganservices: pelangganservices,
	}
}
