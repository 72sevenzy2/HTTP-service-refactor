package routes

import (
	"net/http"

	"github.com/72sevenzy2/http-router/router"

	"github.com/72sevenzy2/golang-API/internal/handler"
	"github.com/72sevenzy2/golang-API/internal/health"
	"github.com/72sevenzy2/golang-API/internal/service"
)

func NewRouter(service service.Greeter) http.Handler {
	r := router.NewRouter()

	r.Handle(http.MethodGet, "/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("response"))
	})

	// usual handlers
	r.Handle(http.MethodPost, "/greet", handler.GreetHandler(service), router.Logger())
	r.Handle(http.MethodGet, "/health", health.HealthHandler(), router.Logger())
	
	return r
}
