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

	tarifServices := services.NewTarifServices(tarifrepository, db, validasi)

	tarifcontroller := controller.NewTarifControllerImpl(tarifServices)

	router := router.NewRouter(tarifcontroller)

	server := http.Server{
		Addr:    "192.168.1.12:8080",
		Handler: middelware.Newmiddelware(router),
	}

	fmt.Println("status ok")
	errsr := server.ListenAndServe()
	helper.Err(errsr)

}
