package main

import (
	_ "github.com/bmizerany/pq"
	"github.com/astaxie/beedb"
	"database/sql"
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
	db_db, err := sql.Open("postgres", "user=adsms dbname=novigeroi password=adsms sslmode=disable")
	db.db = db_db;
	if (err != nil) {
		panic(err.Error())
	}

	// Beedb ORM

	beedb.OnDebug = true

	orm := beedb.New(db.db, "pg")
	db.orm = &orm
}
