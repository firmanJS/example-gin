package config

import (
	"context"
	"fmt"
	"log"
	"time"

	service "firmanjs/example-gin/services/v1"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Connect func to db
func Connect() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbName := "belajar"
	dbHost := "127.0.0.1"
	dbPort := "27017"

	clientOptions := options.Client().ApplyURI("mongodb://" + dbHost + ":" + dbPort)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		print(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var errs = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("cant connect to db", errs)
	} else {
		log.Println("Connected")
	}

	db := client.Database(dbName)
	service.UsersCollections(db)
}
