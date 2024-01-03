package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"shopping-cart/db"
	"shopping-cart/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	mongodbURI := os.Getenv("MONGODB_URI")
	client := db.ConnectDB(mongodbURI)

	defer client.Disconnect(context.Background())

	http.Handle("/", handlers.CorsMiddleware(http.HandlerFunc(handler)))
	http.Handle("/items", handlers.CorsMiddleware(http.HandlerFunc(handlers.GetAllItemsHandler)))
	http.Handle("/cart", handlers.CorsMiddleware(http.HandlerFunc(handlers.NewPostCartItemsHandler(client))))
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
