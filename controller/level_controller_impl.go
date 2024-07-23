package controller

import (
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"golang_listrik/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LevelControllerImpl struct {
	levelServices services.LevelServices
}

// Delete implements LevelController.
func (sv *LevelControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.LevelSearch{}
	helper.Decode_Json(r, &temp)

	sv.levelServices.Delete(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   http.StatusAccepted,
		Status: "accpeth",
		Data:   "data deleted",
	}
	helper.Encode_Json(w, webResponse)
}

// FindAll implements LevelController.
func (sv *LevelControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := sv.levelServices.FIndAll(r.Context())

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   res,
	}

	helper.Encode_Json(w, webResponse)
}

// FindById implements LevelController.
func (sv *LevelControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.LevelSearch{}
	helper.Decode_Json(r, &temp)

	res := sv.levelServices.FindById(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   res,
	}

	helper.Encode_Json(w, webResponse)
}

// Save implements LevelController.
func (sv *LevelControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.LevelRequest{}

	helper.Decode_Json(r, &temp)
	sv.levelServices.Save(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "created",
		Data:   "be ready",
	}
	helper.Encode_Json(w, webResponse)
}

// Update implements LevelController.
func (sv *LevelControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	temp := request.LevelUpdate{}

	helper.Decode_Json(r, &temp)
	sv.levelServices.Update(r.Context(), temp)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   "data Updated",
	}
	helper.Encode_Json(w, webResponse)
}

func NewLevel(levelServices services.LevelServices) LevelController {
	return &LevelControllerImpl{
		levelServices: levelServices,
	}
}
