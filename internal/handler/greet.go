package handler

import (
	"net/http"
	"encoding/json"
	"github.com/72sevenzy2/golang-API/internal/service"
	"github.com/72sevenzy2/golang-API/internal/response"
)

type GreetResponse struct {
	Message string `json:"message"`
	Count int `json:"count"`
}

type GreetRequest struct {
	Name string `json:"name"`
}

func GreetHandler(g service.Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GreetRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid json")
			return
		}

		msg, count, err := g.Greet(req.Name)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		resp := GreetResponse{
			Message: msg,
			Count: count,
		}

		response.Json(w, http.StatusOK, resp)
	}
}