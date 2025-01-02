package main

import (
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connectionString := os.Getenv("DATABASE_URL")
	dbWpo, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// close connection at the end of main
	defer dbWpo.Close()
	// Test the connection to the database
	err = dbWpo.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}
