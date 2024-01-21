package database

import (
	"caching-user-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init(dsn string) (*gorm.DB, error) {
	dbConn, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	if err != nil {
		return nil, err
	}

	dbConn.AutoMigrate(&models.User{})

	db = dbConn
	return db, nil
}

func GetDatabaseConnection() *gorm.DB {
	return db
}
