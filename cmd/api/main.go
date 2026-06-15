package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Bertin-Dreyer/go-social/internal/env"
	"github.com/Bertin-Dreyer/go-social/internal/store"
)

func main() {
	store := store.NewPostgresStorage(nil)

	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))
	fmt.Println("ADDR from env:", os.Getenv("ADDR"))
}
