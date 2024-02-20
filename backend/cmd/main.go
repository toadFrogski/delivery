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
	conn := db.Connect(context.Background())
	if conn == nil {
		log.Fatalf("unable to create connection pool")
	}
	defer conn.Close()

	initHandlers(mux, conn)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func initHandlers(router *http.ServeMux, db *db.Postgres) {
	handlers.InitSampleHandlers(router, db)
}
