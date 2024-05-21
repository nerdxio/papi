package router

import (
	"github.com/nerdxio/chi-demo/internal/handler"
	"github.com/nerdxio/chi-demo/internal/repository/order"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func LoadRouters(rdb *redis.Client) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// sub route
	router.Route("/orders", func(r chi.Router) {
		loadOrderRoutes(rdb, r)
	})

	return router
}

func loadOrderRoutes(rdb *redis.Client, r chi.Router) {
	oh := handler.Order{
		Repo: &order.RedisRepo{
			Client: rdb,
		},
	}
	r.Post("/", oh.Create)
	r.Put("/{id}", oh.UpdateByID)
	r.Delete("/{id}", oh.DeleteByID)
	r.Get("/{id}", oh.GetById)
	r.Get("/", oh.List)
}
