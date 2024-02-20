package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Conn *pgxpool.Pool
}

var (
	pgOnce     sync.Once
	pgInstance *Postgres
)

func Connect(ctx context.Context) *Postgres {
	pgOnce.Do(func() {
		conn, err := pgxpool.New(ctx, getURL())
		if err != nil {
			return
		}

		pgInstance = &Postgres{conn}
	})

	return pgInstance
}

func (p *Postgres) Close() {
	p.Conn.Close()
}

func getURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
}
