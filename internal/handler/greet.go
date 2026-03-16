package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/72sevenzy2/golang-API/internal/service"
	"github.com/72sevenzy2/golang-API/internal/response"
)

type GreetResponse struct {
	Message string `json:"name"`
	Count int `json:"count"`
}

type GreetRequest struct {
	Name string `json:"name"`
}

func GreetHandler(g service.Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GreetRequest

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			json.
		}
	}
}