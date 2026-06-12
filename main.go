package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./jobs.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS jobs (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"payload" TEXT,
		"priority" INTEGER,
		"status" TEXT,
		"max_retries" INTEGER,
		"retries" INTEGER DEFAULT 0
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	// TODO: Implement POST /jobs
	
	// TODO: Implement GET /ws for workers
	
	// TODO: Implement worker heartbeat failure detection

	log.Println("Job Queue Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
