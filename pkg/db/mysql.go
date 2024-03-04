package db

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gostart/pkg/config"
)

var Mysql *gorm.DB

func ConnectMysql() {
	var err error
	p := config.GetString("MYSQL_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GetString("MYSQL_USER"), config.GetString("MYSQL_PASSWORD"), config.GetString("MYSQL_HOST"), port, config.GetString("MYSQL_DATABASE"))

	Mysql, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}
