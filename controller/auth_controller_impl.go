package controller

import (
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	validate *validator.Validate
	Db       *sql.DB
}

// Login implements AuthController.
func (controller *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := r.Context()
	temp := request.LoginRequest{}
	helper.Decode_Json(r, temp)

	tx, _ := controller.Db.Begin()

	err := controller.validate.Struct(temp)
	helper.Err(err)

	if temp.Username[0] != 54 {
		res, e := helper.AdminLogin(ctx, tx, temp.Username, temp.Password)

		if e != nil {
			webResponse := response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "gagal login",
				Data:   e.Error(),
			}
			helper.Encode_Json(w, webResponse)
		}

		webResponse := response.WebResponse{
			Code:   http.StatusOK,
			Status: "ok",
			Data:   res,
		}
		helper.Encode_Json(w, webResponse)
	} else {
		res, e := helper.PelangganLogin(ctx, tx, temp.Username, temp.Password)

		if e != nil {
			webResponse := response.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "gagal login",
				Data:   e.Error(),
			}
			helper.Encode_Json(w, webResponse)
		}

		webResponse := response.WebResponse{
			Code:   http.StatusOK,
			Status: "ok",
			Data:   res,
		}
		helper.Encode_Json(w, webResponse)
	}

}

func NewAuthControllerImpl(validate *validator.Validate, db *sql.DB) AuthController {
	return &AuthControllerImpl{
		validate: validate,
		Db:       db,
	}
}
