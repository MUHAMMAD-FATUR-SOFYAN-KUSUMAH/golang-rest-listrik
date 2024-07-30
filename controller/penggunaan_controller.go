package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PenggunaanController interface {
	FindAllDetail(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	Save(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(writer http.ResponseWriter, r *http.Request, p httprouter.Params)
}
