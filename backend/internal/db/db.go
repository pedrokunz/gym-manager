package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./gym.db")
	if err != nil {
		log.Fatalf("Fatal error opening database: %v", err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		status TEXT,
		joined_at DATETIME
	);
	CREATE TABLE IF NOT EXISTS plans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price REAL,
		duration_months INTEGER
	);
	CREATE TABLE IF NOT EXISTS classes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		trainer TEXT,
		schedule TEXT
	);
	CREATE TABLE IF NOT EXISTS invoices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		member_id INTEGER,
		member_name TEXT,
		amount REAL,
		status TEXT,
		date DATETIME,
		FOREIGN KEY(member_id) REFERENCES members(id)
	);	FOREIGN KEY(member_id) REFERENCES members(id)
	);
	`
	_, err = DB.Exec(createTables)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
}

func GetDB() *sql.DB {
	return DB
}

func ExecuteRaw(query string) (*sql.Rows, error) {
	return DB.Query(query)
}
