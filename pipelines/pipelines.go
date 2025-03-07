package pipelines

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "./timeseries.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table for storing time series data
	createTableSQL := `CREATE TABLE IF NOT EXISTS timeseries (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "timestamp" DATETIME NOT NULL,
        "value" REAL NOT NULL
    );`

	// Create Events Table

	// Create City Table

	// Create WeatherPoll Table

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully")

	// Insert time series data
	insertTimeSeriesData(db, time.Now(), 23.5)
	insertTimeSeriesData(db, time.Now().Add(time.Hour), 24.0)
	insertTimeSeriesData(db, time.Now().Add(2*time.Hour), 22.8)

	// Insert Faux Data

	// Insert 69Racing Event 6/9/2025 Thunderhill West Willows, CA

	fmt.Println("Data inserted successfully")
}

func insertTimeSeriesData(db *sql.DB, timestamp time.Time, value float64) {
	insertSQL := `INSERT INTO timeseries (timestamp, value) VALUES (?, ?)`
	_, err := db.Exec(insertSQL, timestamp, value)
	if err != nil {
		log.Fatal(err)
	}
}
