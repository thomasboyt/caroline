package store

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func New() *Store {
	dsn := os.Getenv("DATABASE_URL")
	dbx, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	db := &Store{dbx}
	return db
}
