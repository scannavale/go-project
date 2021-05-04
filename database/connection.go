package database

import (
	"go-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=X360_l123 dbname=go_project port=5432 TimeZone=Asia/Shanghai"
	connection, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		log.Fatal("Unable to connect to the database")
	}

	DB = connection

	if migrateErr := connection.AutoMigrate(&models.Post{}); migrateErr != nil {
		log.Fatal("Unable to migrate the database")
	}
}
