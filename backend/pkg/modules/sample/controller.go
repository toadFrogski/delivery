package sample

import (
	"delivery/pkg/utils/db"
	"net/http"
)

type SampleController struct {
	sampleSvc *sampleService
}

func NewSampleController(db *db.Postgres) *SampleController {
	sampleRepo := &sampleRepository{db: db.Conn}
	sampleService := &sampleService{sampleRepo: sampleRepo}
	return &SampleController{sampleSvc: sampleService}
}

func (sc *SampleController) Sample(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'test': 'test'}"))
}
