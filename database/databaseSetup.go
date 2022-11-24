package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	//connecting to local mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	//Set context with connection timeOut for secure connection
	dbContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(dbContext)
	if err != nil {
		log.Println("Connection Failed")
		log.Fatal(err)
		return nil
	}

	fmt.Println("Successfuly Connected")
	return client
}

// creating a pointer to a mongodb struct
var Client *mongo.Client = DBSet()

// specifying collection for user and product data
func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("goShop").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("goShop").Collection(collectionName)
	return collection
}
