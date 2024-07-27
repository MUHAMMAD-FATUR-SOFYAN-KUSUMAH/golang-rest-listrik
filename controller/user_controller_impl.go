package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type userControllerImpl struct {
	userService services.UserServices
}

// Delete implements UserController.
func (controller *userControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.UserSearch{}
	helper.Decode_Json(r, &temp)

	controller.userService.Delete(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeth",
		Data:   "data deleted",
	}
	helper.Encode_Json(w, webResponse)
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
func (controller *userControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.UserSearch{}

	helper.Decode_Json(r, &temp)

	res := controller.userService.FindById(r.Context(), temp.Id)

	response := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   res,
	}

	helper.Encode_Json(w, response)
}

// Save implements UserController.
func (controller *userControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.AddUser{}
	helper.Decode_Json(r, &temp)

	controller.userService.Save(r.Context(), temp)

	res := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeted",
		Data:   "data saved",
	}

	helper.Encode_Json(w, res)

}

// Update implements UserController.
func (controller *userControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.UserUpdate{}

	helper.Decode_Json(r, &temp)

	controller.userService.Update(r.Context(), temp)

	res := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeted",
		Data:   "data updated",
	}

	helper.Encode_Json(w, res)
}

func NewUserController(userService services.UserServices) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}
