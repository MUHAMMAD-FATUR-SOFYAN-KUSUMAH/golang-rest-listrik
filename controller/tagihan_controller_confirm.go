package controller

import (
	"database/sql"
	"fmt"
	"golang_listrik/helper"
	"golang_listrik/model/request"
	"golang_listrik/model/response"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	storage_go "github.com/supabase-community/storage-go"
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
	sql := "SELECT t.id_tagihan, pg.bulan , pg.tahun, t.status, t.jumlah_meter FROM public.tagihan as t JOIN public.penggunaan as pg ON pg.pelanggan = t.pelanggan WHERE t.pelanggan = $1"

	row, err := tx.Query(sql, temp.Id)

	helper.Err(err)
	var model []response.TagihanResponse

	for row.Next() {
		var temp_int int
		var gory response.TagihanResponse
		err := row.Scan(&gory.Id_tagihan, &gory.Bulan, &gory.Tahun, &temp_int, &gory.Penggunaan_Kwh)
		helper.Err(err)
		result := helper.Int_StatusToString(temp_int)
		gory.Status = result
		model = append(model, gory)
	}

	helper.Encode_Json(w, model)
}

func (c *TagihanController) UploadProfTagihan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	env, _ := godotenv.Read(".env")
	tx, _ := c.db.Begin()

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}

	var (
		url       = env["Url"]
		token     = env["Token"]
		extension = filepath.Ext(header.Filename)
	)

	switch extension {
	case ".jpg":
		extension = "jpg"
	case ".jpeg":
		extension = "jpeg"
	case ".png":
		extension = "png"
	default:
		webresponse := response.WebResponse{
			Code:   http.StatusNotAcceptable,
			Status: "kesalahan format",
			Data:   "File not uploaded",
		}

		helper.Encode_Json(w, webresponse)
		return
	}

	er := r.ParseMultipartForm(10 << 20) // 10 MB
	if er != nil {
		http.Error(w, "must be less than 10 MB", http.StatusInternalServerError)
		return
	}

	id := r.FormValue("id")
	id_tagihan, err_convert := strconv.Atoi(id)

	helper.Err(err_convert)
	//upload to databae
	helper.UploadImageSql(r.Context(), tx, header.Filename, id_tagihan)
	// create client to supabase storage
	cc := storage_go.NewClient(url, token, nil)
	//content-type
	str := fmt.Sprintf("image/%s", extension)

	_, e := cc.UploadFile("Lolking", header.Filename, file, storage_go.FileOptions{
		ContentType: &str,
	})

	if e != nil {
		webresponse := response.WebResponse{
			Code:   http.StatusNotAcceptable,
			Status: "Not Acceptable",
			Data:   "supabase upload error",
		}

		helper.Encode_Json(w, webresponse)
		return
	}

	defer func() {
		file.Close()
		helper.Tx(tx)
	}()

	webresponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   "File uploaded successfully",
	}

	helper.Encode_Json(w, webresponse)
}
