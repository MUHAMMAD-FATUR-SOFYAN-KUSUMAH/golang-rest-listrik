package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type userControllerImpl struct {
	userService services.UserServices
}

// Delete implements UserController.
func (*userControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements UserController.
func (controller *userControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := controller.userService.FindAll(r.Context())

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   res,
	}
	helper.Encode_Json(w, webResponse)

}

// FindById implements UserController.
func (*userControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Save implements UserController.
func (*userControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Update implements UserController.
func (*userControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

func NewUserController(userService services.UserServices) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}
