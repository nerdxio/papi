package main

import (
	"context"
	"log"

	"github.com/nerdxio/chi-demo/app"
)

func main() {

	app := app.New()

	err := app.Start(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}
