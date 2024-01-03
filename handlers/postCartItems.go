package handlers

import (
	"encoding/json"
	"net/http"
	"shopping-cart/functions"
	"shopping-cart/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewPostCartItemsHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cartItems []models.CartData

		err := json.NewDecoder(r.Body).Decode(&cartItems)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = functions.InsertCartItems(cartItems, client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cart items added successfully"))
	}
}
