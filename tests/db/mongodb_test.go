package db

import (
	"context"
	"fmt"
	"gostart/pkg/config"
	db "gostart/pkg/db"
	"gostart/pkg/schema/mongodb"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestWrite(t *testing.T) {
	config.Load("../../.env")
	db.ConnectMongoDB()

	coll := db.Mdb().Collection(db.Log)

	log := &mongodb.Log{}
	log.ID = primitive.NewObjectID()
	log.CreatedAt = time.Now()
	fmt.Println(log.ID)
	result, err := coll.InsertOne(context.Background(), log)
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result.InsertedID != log.ID {
		t.Errorf("Expected %v, got %v", log.ID, result.InsertedID)
	}
}
