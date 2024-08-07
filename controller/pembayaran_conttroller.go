package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PembayaranController interface {
	FindAllKonfirmasi(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllDetails(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
}
