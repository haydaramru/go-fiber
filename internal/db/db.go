package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

type Collection string

const (
	ProductsCollection Collection = "products"
	// UsersCollection    Collection = "users"
)

const (
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.85fxdxq.mongodb.net/", username, password)

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connectionString)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client

		clientInstanceError = err
	})
	return clientInstance, clientInstanceError
}

// 💡💡💡
// * is reference the value (it is the pointer, probably)
// & is point the block of memory that stores this object
