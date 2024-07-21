package db

import (
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", "gophkeeper.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&entities.User{}, &entities.Data{})

	log.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {
	return DB
}
