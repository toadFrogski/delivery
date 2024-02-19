package main

import (
	"context"
	"delivery/pkg/handlers"
	"delivery/pkg/utils/db"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	db, err := db.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	initHandlers(mux, db)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func initHandlers(router *http.ServeMux, db *db.Postgres) {
	handlers.InitSampleHandlers(router, db)
}
