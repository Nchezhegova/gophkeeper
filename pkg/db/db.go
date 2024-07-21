package db

import (
	"github.com/Nchezhegova/gophkeeper/internal/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gophkeeper.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Data{})

	log.Println("Database connection successfully established")
	return db
}
