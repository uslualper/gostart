package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Level     int                `bson:"level"`
	Type      string             `bson:"type"`
	Owner     string             `bson:"owner"`
	Context   interface{}        `bson:"context"`
	CreatedAt time.Time          `bson:"created_at"`
}
