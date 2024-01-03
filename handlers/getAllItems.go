package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"shopping-cart/db"
	"shopping-cart/functions"

	"github.com/joho/godotenv"
)

func GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodbURI := os.Getenv("MONGODB_URI")
	client := db.ConnectDB(mongodbURI)

	defer client.Disconnect(context.Background())

	items, err := functions.GetAllCartItems(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
