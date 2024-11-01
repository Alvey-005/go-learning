package main

import (
	"CRUD-Project/internal/db"
	"CRUD-Project/internal/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Creating a context for connecting to the database
	ctx := context.Background()

	// Constructing the connection string from environment variables
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// Initializing the database connection
	dbInstance, err := pg.NewPG(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer dbInstance.Close() // Ensure the connection is closed when done

	// Test the connection
	if err := dbInstance.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection established successfully.")
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api", checkStatus).Methods("GET")

	log.Println("Starting server on :5000")

	http.ListenAndServe(":5000", r)
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("wokring properly")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
