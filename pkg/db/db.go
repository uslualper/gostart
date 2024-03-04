package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func SetupDB() {
	ConnectMysql()
	ConnectMongoDB()
}

func Sql() *gorm.DB {
	return Mysql
}

func Mdb() *mongo.Database {
	return MongoDB
}
