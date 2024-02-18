package main

import (
	"delivery/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handlers.InitHandlers(mux)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
