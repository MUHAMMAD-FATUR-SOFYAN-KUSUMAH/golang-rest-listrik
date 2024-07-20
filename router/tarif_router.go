package router

import (
	"golang_listrik/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tarif controller.TarifController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tarif", tarif.FindAll)

	return router
}
