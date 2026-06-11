package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Bertin-Dreyer/go-social/internal/env"
)

func main() {
	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
	fmt.Println("ADDR from env:", os.Getenv("ADDR"))
}
