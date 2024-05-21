package main

import (
	"context"
	"github.com/nerdxio/chi-demo/internal/app"
	"github.com/nerdxio/chi-demo/internal/router"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := app.New()
	app.Router = router.LoadRouters(app.Rdb)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
