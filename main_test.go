package main

import (
	"testing"
)

func TestDatabaseInitialization(t *testing.T) {
	initDB()
	if db == nil {
		t.Fatal("Database connection is nil")
	}
	
	// Verify jobs table exists
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='jobs';")
	if err != nil {
		t.Fatalf("Failed to query database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Fatal("Table 'jobs' was not created by initDB()")
	}
}
