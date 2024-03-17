package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/nerdxio/chi-demo/app"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
