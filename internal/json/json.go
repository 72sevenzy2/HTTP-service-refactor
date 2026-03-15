package json

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, message string) {
	res := map[string]string{
		"error": message,
	}

	Json(w, status, res)
}