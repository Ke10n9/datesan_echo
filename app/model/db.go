package model

import (
	"log"

	"app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// connect DB
	conf := config.Config
	dsn := conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:root@unix(http://localhost:8888/Applications/MAMP/tmp/mysql/mysql.sock)/datesan_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // NOTE: using = to global var , := is only this function
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
}