package logger

import (
	"context"
	"gostart/pkg/db"
	logSchema "gostart/pkg/schema/logger"
	mongoSchema "gostart/pkg/schema/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongo struct{}

func (m *Mongo) Save(owner OwnerType, logContext logSchema.Context, typ LogType, level int) {
	log := mongoSchema.Log{
		ID:        primitive.NewObjectID(),
		Owner:     string(owner),
		Context:   logContext,
		Type:      string(typ),
		Level:     level,
		CreatedAt: time.Now(),
	}

	_, err := db.Mdb().Collection(db.Log).InsertOne(context.TODO(), log)
	if err != nil {
		panic(err)
	}
}
