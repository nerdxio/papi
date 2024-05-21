package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	Router http.Handler
	Rdb    *redis.Client
}

func New() *App {
	return &App{
		Rdb: redis.NewClient(&redis.Options{}),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.Router,
	}

	err := a.Rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("fail to connect to redis client %w", err)
	}

	defer func() {
		if err := a.Rdb.Close(); err != nil {
			log.Println("Failed to close redis", err)
		}
	}()

	log.Println("Starting the Server on port ", server.Addr)

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("fail to start the server %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
