package logger

import (
	"context"
	"gostart/pkg/db"
	mongoSchema "gostart/pkg/schema/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongo struct{}

func (m *Mongo) Save(owner OwnerType, logContext Context, typ LogType, level int) {
	log := mongoSchema.Log{
		ID:        primitive.NewObjectID(),
		Owner:     string(owner),
		Type:      string(typ),
		Level:     level,
		CreatedAt: time.Now(),
	}

	_, err := db.Mdb().Collection(db.Log).InsertOne(context.TODO(), log)
	if err != nil {
		panic(err)
	}
}
