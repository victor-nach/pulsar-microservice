package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - initialize connection to mongo db
func Connect() (*mongo.Collection, error) {
	goENV, ok := os.LookupEnv("GO_ENV")
	if !ok {
		goENV = "LOCAL"
	}
	var dbURL, dbName string
	var URLOk, nameOk bool
	if goENV == "TEST" {
		dbURL, URLOk = os.LookupEnv("DB_URL_TEST")
		dbName, nameOk = os.LookupEnv("DB_NAME_TEST")
	} else {
		dbURL, URLOk = os.LookupEnv("DB_URL")
		dbName, nameOk = os.LookupEnv("DB_NAME")
	}

	if !URLOk || !nameOk {
		log.Fatal("db env not set")
	}
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("Connected to the db...")
	collection := client.Database(dbName).Collection("users")
	return collection, nil
}
