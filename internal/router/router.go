package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/72sevenzy2/golang-API/internal/handler"
	"github.com/72sevenzy2/golang-API/internal/service"
	"github.com/72sevenzy2/golang-API/internal/health"
)

func NewRouter(service service.Greeter) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/greet", handler.GreetHandler(service))
	r.Get("/health", health.HealthHandler())

	return r
}