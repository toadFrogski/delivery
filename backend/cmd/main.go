package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, _ := json.Marshal("hello world")
		io.Writer.Write(w, response)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
