package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dannyoka/memorize-it-api/internal/controllers"
	"github.com/dannyoka/memorize-it-api/internal/repositories"
	"github.com/dannyoka/memorize-it-api/internal/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func loadEnv() {
	if os.Getenv("ENV") == "development" {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}
}

func createMongoClient() *mongo.Client {
	client, err := mongo.Connect(
		options.Client().ApplyURI(os.Getenv("MONGODB_URI")),
	)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	fmt.Println("Now starting memorize-it-api")
	loadEnv()
	client := createMongoClient()
	defer client.Disconnect(nil)
	db := client.Database("memorize-it")
	entryRepository := repositories.NewEntryRepository(db)
	entryService := services.NewEntryService(*entryRepository)
	entryController := controllers.NewEntryController(*entryService)
	http.HandleFunc("/entries", entryController.HandleGetEntries)
	http.ListenAndServe(":8080", nil)
}
