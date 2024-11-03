package handlers

import (
	"CRUD-Project/internal/db"
	"CRUD-Project/internal/models"
	"context"
	"encoding/json"

	// "io"
	"log"
	"net/http"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	// Decode the JSON payload
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("Error decoding JSON: %v", err) // Log the error for debugging
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the pg singleton instance
	pgInstance, err := pg.NewPG(context.Background(), "your_connection_string_here")
	if err != nil {
		log.Printf("Failed to get DB instance: %v", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	// Execute the query to insert the new item
	query := "INSERT INTO items (name, price) VALUES ($1, $2) RETURNING id"
	err = pgInstance.GetDB().QueryRow(context.Background(), query, item.Name, item.Price).Scan(&item.ID)
	if err != nil {
		log.Printf("Failed to create item: %v", err) // Log the error for debugging
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	// Respond with the created item
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
