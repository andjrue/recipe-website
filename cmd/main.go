package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Printf("Error loading env: %v", envErr)
	}

	fmt.Println("env loaded")

	db, err := sql.Open("pgx", os.Getenv("PG_DSN"))
	fmt.Println("sql.open")
	if err != nil {
		fmt.Printf("Unable to connect to Database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	fmt.Println("Pinging db")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to DB")

	server := newApiServer(":3000")
	server.Run()

}
