package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nerdxio/chi-demo/handler"
	"github.com/nerdxio/chi-demo/repository/order"
)

func (a *App) loadRouters() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// sub route
	router.Route("/orders", a.loadOrderRoutes)
	a.router = router
}

func (a *App) loadOrderRoutes(r chi.Router) {
	oh := handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}
	r.Post("/", oh.Create)
	r.Put("/{id}", oh.UpdateById)
	r.Delete("/{id}", oh.DeleteById)
	r.Get("/{id}", oh.GetById)
	r.Get("/", oh.List)
}
