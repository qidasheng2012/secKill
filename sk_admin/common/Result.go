package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func ResultOk(w http.ResponseWriter, data string) {
	io.WriteString(w, data)
}

func ResultFail(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func ResultJsonOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}
