package functions

import (
	"context"
	"shopping-cart/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertCartItems(cartItems []models.CartData, client *mongo.Client) error {
	collection := client.Database("shopping-cart").Collection("added-items")

	var documents []interface{}
	for _, item := range cartItems {
		documents = append(documents, item)
	}

	_, err := collection.InsertMany(context.Background(), documents)
	return err
}
