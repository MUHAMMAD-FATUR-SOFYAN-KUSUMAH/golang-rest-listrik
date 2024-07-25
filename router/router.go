package router

import (
	"golang_listrik/controller"
	"golang_listrik/helper"
	"golang_listrik/middelware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tarif controller.TarifController, level controller.LevelController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/tarif", helper.WrapMiddelware(httprouter.Handle(tarif.FindAll), middelware.LoggingMiddleware))
	router.POST("/api/tarif", tarif.Save)
	router.DELETE("/api/tarif", tarif.Deleted)
	router.PATCH("/api/tarif", tarif.Updated)
	router.GET("/api/tarifs", tarif.FindById)

	router.GET("/api/level", level.FindAll)
	router.POST("/api/level", level.Save)
	router.DELETE("/api/level", level.Delete)
	router.PATCH("/api/level", level.Update)
	router.GET("/api/levels", level.FindById)

	return router
}
