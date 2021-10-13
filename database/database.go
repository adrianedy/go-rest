package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection *mongo.Client
var database string = "mflix"

func GetConnection() *mongo.Client {
	if connection == nil {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		connection, _ = mongo.Connect(context.TODO(), clientOptions)
	}

	return connection
}

func Collection(input string) *mongo.Collection {
	return GetConnection().Database(database).Collection(input)
}
