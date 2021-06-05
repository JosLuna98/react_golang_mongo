package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"server/controllers"
	"server/repository"
	"server/router"
)

func main() {
	tr := repository.NewTaskRepository(getCollectionInstance())
	tc := controllers.NewTaskController(tr)

	r := router.Router(tc)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getCollectionInstance() *mongo.Collection {
	err := godotenv.Load(".env.dev")

	if err != nil {
		log.Println("Production mode")
	}

	connectionString := os.Getenv("DB_URI")

	dbName := os.Getenv("DB_NAME")

	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var collection *mongo.Collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")

	return collection
}
