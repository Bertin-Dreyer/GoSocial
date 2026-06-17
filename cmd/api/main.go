package main

import (
	"log"

	"github.com/Bertin-Dreyer/go-social/internal/db"
	"github.com/Bertin-Dreyer/go-social/internal/env"
	"github.com/Bertin-Dreyer/go-social/internal/store"
)

func main() {
	addr := env.GetString("ADDR", ":8080")
	dbAddr := env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable")
	maxOpenConns := env.GetInt("DB_MAX_OPEN_CONNS", 30)
	maxIdleConns := env.GetInt("DB_MAX_IDLE_CONNS", 30)
	maxIdleTime := env.GetString("DB_MAX_IDLE_TIME", "15min")

	database, err := db.New(dbAddr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	pgStore := store.NewPostgresStorage(database)

	app := &application{
		config: config{
			addr: addr,
			db: dbConfig{
				addr: dbAddr,
				maxOpenConns: maxOpenConns,
				maxIdleConns: maxIdleConns,
				maxIdleTime: maxIdleTime,
			},
		},
		store: pgStore,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}