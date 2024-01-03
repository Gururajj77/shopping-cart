package functions

import (
	"context"
	"log"
	"shopping-cart/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllCartItems(client *mongo.Client) ([]models.CartData, error) {
	var cartItems []models.CartData

	collection := client.Database("shopping-cart").Collection("shop-items")
	cursor, err := collection.Find(context.Background(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var cartItem models.CartData
		err := cursor.Decode(&cartItem)
		if err != nil {
			log.Fatal(err)
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return cartItems, nil
}
