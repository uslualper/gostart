package db

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-start/config"
)

var Mysql *gorm.DB

func ConnectMysql() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))

	Mysql, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}
