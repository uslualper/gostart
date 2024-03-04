package db

import (
	"context"
	"fmt"
	"gostart/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

const (
	User = "users"
	Log  = "logs"
)

func ConnectMongoDB() {
	name := config.GetString("MONGODB_NAME")
	url := config.GetString("MONGODB_URL")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	MongoDB = client.Database(name)

	fmt.Println("Connection Opened to MongoDB")
}
