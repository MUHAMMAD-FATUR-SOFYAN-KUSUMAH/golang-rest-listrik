package main

import (
	"fmt"
	"golang_listrik/config"
	"golang_listrik/controller"
	"golang_listrik/helper"
	"golang_listrik/middelware"
	"golang_listrik/repository"
	"golang_listrik/router"
	"golang_listrik/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.NewDb()
	validasi := validator.New()

	tarifrepository := repository.NewTarifRepository()
	levelreposiotry := repository.NewLevelRepository()

	tarifServices := services.NewTarifServices(tarifrepository, db, validasi)
	levelservices := services.NewLevelServices(levelreposiotry, db, validasi)

	tarifcontroller := controller.NewTarifControllerImpl(tarifServices)
	levelcontroller := controller.NewLevel(levelservices)

	routers := router.NewRouter(tarifcontroller, levelcontroller)

	server := http.Server{
		Addr:    "192.168.1.12:8080",
		Handler: middelware.Newmiddelware(routers),
	}

	fmt.Println("status ok")
	errsr := server.ListenAndServe()
	helper.Err(errsr)

}
