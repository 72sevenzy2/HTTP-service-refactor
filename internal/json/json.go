package json

import (
	"encoding/json"
	"net/http"
)

func json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func error(w http.ResponseWriter, status int, message string) {
	res := map[string]string{
		"error": message,
	}

	json(w, status, res)
}