package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nerdxio/chi-demo/handler"
)

func loadRouters() *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// sub route
	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(r chi.Router) {
	oh := handler.Order{}
	r.Post("/", oh.Create)
	r.Put("/{id}", oh.UpdateById)
	r.Delete("/{id}", oh.DeleteById)
	r.Get("/{id}", oh.GetById)
	r.Get("/", oh.List)
}
