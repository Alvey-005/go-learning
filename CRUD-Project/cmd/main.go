package main

import (
	"CRUD-Project/internal/db"
	"CRUD-Project/internal/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	// defer dbInstance.Close() // Ensure the connection is closed when done

	// Test the connection
	if err := dbInstance.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection established successfully.")
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api", checkStatus).Methods("GET")

	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")

	http.HandleFunc("/init-db", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		connString := "your_connection_string" // Replace with your actual connection string
		db, err := pg.NewPG(ctx, connString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to initialize DB: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "DB initialized successfully: %+v", db)
	})

	log.Println("Starting server on :5000")

	http.ListenAndServe(":5000", r)
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("wokring properly")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
