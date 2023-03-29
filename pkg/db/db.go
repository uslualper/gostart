package db

import "gorm.io/gorm"

func SetupDB() {
	ConnectMysql()
}

func Sql() *gorm.DB {
	return Mysql
}
