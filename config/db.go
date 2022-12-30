package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:mysql-local@tcp(127.0.0.1)/belajar-grpc"))
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	return db
}
