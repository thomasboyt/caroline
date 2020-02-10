package store

import (
	"log"

	"github.com/DATA-DOG/go-txdb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func New(dsn string) *Store {
	dbx, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	db := &Store{dbx}
	return db
}

func CreateTestStore() *Store {
	txdb.Register("pgx", "postgres", "postgres://postgres@localhost:5433/jambuds_test?sslmode=disable")
	dbx, err := sqlx.Connect("pgx", "CreateTestStore")

	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	db := &Store{dbx}
	return db
}
