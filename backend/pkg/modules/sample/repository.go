package sample

import "github.com/jackc/pgx/v5/pgxpool"

type sampleRepository struct {
	db *pgxpool.Pool
}
