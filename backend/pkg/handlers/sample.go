package handlers

import (
	"delivery/pkg/modules/sample"
	"delivery/pkg/utils/db"
	"net/http"
)

func InitSampleHandlers(router *http.ServeMux, db *db.Postgres) {
	sample := sample.NewSampleController(db)

	router.HandleFunc("/", sample.Sample)
}
