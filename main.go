package main

import (
	"context"

	"github.com/wendermrn/barber-shop-go-api/pkg/api"
)

//var seed = flag.Bool("seed", true, "Pre populate database")

func main() {
	/*if *seed {
		log.Info("Init Seed")
		database.Seed()
	}*/
	ctx := context.Background()
	api.Run(ctx)
}
