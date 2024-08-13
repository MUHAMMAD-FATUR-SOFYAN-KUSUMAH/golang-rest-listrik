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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "default url")
}

func main() {
	db := config.NewDb()
	validasi := validator.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/auth", rootHandler)

	tarifrepository := repository.NewTarifRepository()
	levelreposiotry := repository.NewLevelRepository()
	userrepository := repository.NewUserRepository()
	pelanggan := repository.NewPelangganRepsitoryImpl()
	Penggunaan := repository.NewPenggunaanRepositoryImpl()
	Pembayaran := repository.NewPembayaranRepository()

	tarifServices := services.NewTarifServices(tarifrepository, db, validasi)
	levelservices := services.NewLevelServices(levelreposiotry, db, validasi)
	userservices := services.NewUserServicesImpl(db, userrepository, validasi)
	pelangganservices := services.NewPelangganServicesImpl(pelanggan, db, validasi)
	penggunaanservices := services.NewPenggunaanServices(validasi, db, Penggunaan)
	pembayaranservices := services.NewPembayaranServices8(db, Pembayaran)

	tarifcontroller := controller.NewTarifControllerImpl(tarifServices)
	levelcontroller := controller.NewLevel(levelservices)
	usercontroller := controller.NewUserController(userservices)
	pelanggancontroller := controller.NewPelangganControllerImpl(pelangganservices)
	penggunaancontroller := controller.NewPenggunaanController(penggunaanservices)
	tagihancontroller := controller.NewTagihanController(db)
	PembayaranController := controller.NewPembayaranController(pembayaranservices)
	LoginController := controller.NewAuthControllerImpl(validasi, db)

	routers := router.NewRouter(tarifcontroller, levelcontroller, usercontroller, pelanggancontroller, penggunaancontroller, tagihancontroller, PembayaranController)
	AuthRouter := router.NewRouterAuthLogin(LoginController)

	mux.Handle("/api/", http.StripPrefix("/api", routers))
	mux.Handle("/auth/", http.StripPrefix("/auth", AuthRouter))
	// middelware add
	initMiddelwarevalidasi := &middelware.Authmiddelware{
		Handler: mux,
	}

	initMIddelwareCors := &middelware.CorsMiddelware{
		Handler: initMiddelwarevalidasi,
	}

	server := http.Server{
		Addr:    "192.168.1.2:8080",
		Handler: initMIddelwareCors,
	}

	fmt.Println("status ok")
	errsr := server.ListenAndServe()
	helper.Err(errsr)
}
