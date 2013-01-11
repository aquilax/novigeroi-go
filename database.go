package main

import (
	_ "github.com/bmizerany/pq"
	"database/sql"
	"log"
)

type Database struct {
	connected bool
	db *sql.DB
}


func NewDatabase () *Database {
	return &Database{false, nil}
}

func (db *Database) Connect () {
	db_db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	db.db = db_db;
	if (err != nil) {
		log.Fatal(err)
	}
	log.Print("123")
}
