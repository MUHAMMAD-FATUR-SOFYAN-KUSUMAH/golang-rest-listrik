package helper

import (
	"encoding/json"
	"net/http"
)

func Decode_Json(req *http.Request, result any) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	Err(err)
}

func Encode_Json(w http.ResponseWriter, result any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)
	Err(err)
}
