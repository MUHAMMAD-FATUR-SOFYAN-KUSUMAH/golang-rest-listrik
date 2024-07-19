package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TarifController interface {
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Updated(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Deleted(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
