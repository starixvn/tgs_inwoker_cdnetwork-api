package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Replace the connection string with your PostgreSQL connection string
	connStr := "postgres://username:password@localhost/dbname?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Execute a query to select data from the agent_domain table
	rows, err := db.Query("SELECT * FROM agent_domain")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var (
			id     int
			domain string
		)
		// Scan the values from the row into variables
		err := rows.Scan(&id, &domain)
		if err != nil {
			panic(err)
		}
		// Print the values
		fmt.Printf("ID: %d, Domain: %s\n", id, domain)
	}
	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
