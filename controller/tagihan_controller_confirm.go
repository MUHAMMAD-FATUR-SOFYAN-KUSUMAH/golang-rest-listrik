package controller

import (
	"database/sql"
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TagihanController struct {
	db *sql.DB
}

func NewTagihanController(db *sql.DB) *TagihanController {
	return &TagihanController{db: db}
}

// FindAll implements TagihanController
func (c *TagihanController) FindPelangganTagihan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var temp request.PelangganRequest

	helper.Decode_Json(r, &temp)
	tx, _ := c.db.Begin()
	sql := "SELECT pg.bulan , pg.tahun, t.status, t.jumlah_meter FROM public.tagihan as t JOIN public.penggunaan as pg ON pg.pelanggan = t.pelanggan WHERE t.pelanggan = $1"

	row, err := tx.Query(sql, temp.Id)

	helper.Err(err)
	var model []response.TagihanResponse

	for row.Next() {
		var temp_int int
		var gory response.TagihanResponse
		err := row.Scan(&gory.Bulan, &gory.Tahun, &temp_int, &gory.Penggunaan_Kwh)
		helper.Err(err)
		result := helper.Int_StatusToString(temp_int)
		gory.Status = result
		model = append(model, gory)
	}

	helper.Encode_Json(w, model)
}
