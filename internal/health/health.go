package health

import (
	"github.com/72sevenzy2/golang-API/internal/response"
	"net/http"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ignore requests which arent GET
		if r.Method != http.MethodGet {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
			return
		}

		resp := map[string]string{
			"message": "fully working API",
		}

		response.Json(w, http.StatusOK, resp)
	}
}
