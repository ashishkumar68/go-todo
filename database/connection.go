package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetManager() (*gorm.DB, error) {
	if connection != nil {
		return connection, nil
	}
	conn, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection = conn

	return connection, nil
}
