package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func newDB() (*Storage, error) {
	connStr := "user=postgres dbname=postgres password=recipe sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}
