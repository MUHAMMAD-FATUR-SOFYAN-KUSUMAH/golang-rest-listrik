package router

import (
	"golang_listrik/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouterAuthLogin(auth controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", auth.Login)

	return router
}
