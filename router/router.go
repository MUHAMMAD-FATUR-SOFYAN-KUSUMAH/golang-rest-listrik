package router

import (
	"golang_listrik/controller"
	"golang_listrik/helper"
	"golang_listrik/middelware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tarif controller.TarifController, level controller.LevelController, user controller.UserController, pelanggan controller.PelangganController, penggunaan controller.PenggunaanController) *httprouter.Router {
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

	router.GET("/api/user", user.FindAll)
	router.GET("/api/users", user.FindById)
	router.POST("/api/user", user.Save)
	router.DELETE("/api/user", user.Delete)
	router.PATCH("/api/user", user.Update)

	router.GET("/api/pelanggan", pelanggan.FindAll)
	router.GET("/api/pelanggans", pelanggan.FindById)
	router.POST("/api/pelanggan", pelanggan.Save)
	router.DELETE("/api/pelanggan", pelanggan.Delete)
	router.PATCH("/api/pelanggan", pelanggan.Update)

	router.GET("/api/penggunaan", penggunaan.FindAll)
	router.GET("/api/penggunaan/detail", penggunaan.FindAllDetail)
	router.POST("/api/penggunaan", penggunaan.Save)
	router.DELETE("/api/penggunaan", penggunaan.Delete)
	router.PATCH("/api/penggunaan", penggunaan.Update)

	return router
}
