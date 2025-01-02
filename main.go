package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	db, err := sqlx.Connect("postgres", "user=postgres dbname=wpodev sslmode=disable password=Becarefulwithme host=192.168.61.83")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}
