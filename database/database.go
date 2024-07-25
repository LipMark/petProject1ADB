package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/LipMark/ToDo-Go/util"
)

func CreateDB() (*mongo.Client, error) {
	fmt.Println("trying to connect to db")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	util.CheckError(err)

	err = client.Ping(ctx, readpref.Primary())
	util.CheckError(err)

	fmt.Println("Connected to the MongoDB")
	return client, nil
}
