package router

import (
	"golang_listrik/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tarif controller.TarifController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tarif", tarif.FindAll)
	router.POST("/api/tarif", tarif.Save)
	router.DELETE("/api/tarif", tarif.Deleted)

	return router
}
