package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PembayaranControllerImpl struct {
	pembayaranService services.PembayaranServices
}

// Delete implements PembayaranController.
func (controller *PembayaranControllerImpl) Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var temp request.LevelSearch
	helper.Decode_Json(r, &temp)
	controller.pembayaranService.Delete(r.Context(), temp)

	webrespons := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "Pembayaran ditolak",
		Data:   "mohon di ulangi kembali",
	}

	helper.Encode_Json(writer, webrespons)
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
