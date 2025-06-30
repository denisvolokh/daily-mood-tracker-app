package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB(filepath string) *sql.DB {
	database, err := sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTable(database)
	return database
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS mood_entries (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			mood TEXT NOT NULL,
			note TEXT,
			date TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("createTable:", err)
	}
}
