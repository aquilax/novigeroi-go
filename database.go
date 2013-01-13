package main

import (
	_ "github.com/bmizerany/pq"
	"github.com/astaxie/beedb"
	"database/sql"
	"log"
)

type Database struct {
	connected bool
	db *sql.DB
	orm *beedb.Model
}


func NewDatabase () *Database {
	return &Database{false, nil, nil}
}

func (db *Database) Connect () {
	db_db, err := sql.Open("postgres", "user=adsms dbname=novigeroi password=password sslmode=disable")
	db.db = db_db;
	if (err != nil) {
		log.Fatal(err)
	}

	// Beedb ORM

	beedb.OnDebug = true

	orm := beedb.New(db.db, "pg")
	db.orm = &orm
}

func (db *Database) Orm () beedb.Model {
	return *db.orm
}
