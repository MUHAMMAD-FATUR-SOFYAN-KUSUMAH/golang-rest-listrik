package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PelangganController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Save(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
}
