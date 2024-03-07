package router

import (
	"xpends/webapi/handlers"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.Get)
	return router
}
