package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PembayaranControllerImpl struct {
	pembayaranService services.PembayaranServices
}

// Delete implements PembayaranController.
func (*PembayaranControllerImpl) Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindAllDetails implements PembayaranController.
func (controller *PembayaranControllerImpl) FindAllDetails(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := controller.pembayaranService.FindAllDetails(r.Context())

	webrespons := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   result,
	}

	helper.Encode_Json(writer, webrespons)
}

// FindAllKonfirmasi implements PembayaranController.
func (controller *PembayaranControllerImpl) FindAllKonfirmasi(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := controller.pembayaranService.FindAllKonfrimasi(r.Context())

	webrespons := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   result,
	}

	helper.Encode_Json(writer, webrespons)
}

func NewPembayaranController(pembayaranService services.PembayaranServices) PembayaranController {
	return &PembayaranControllerImpl{pembayaranService: pembayaranService}
}
