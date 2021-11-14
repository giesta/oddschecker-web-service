package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "odds_checker.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Odd{})

	DB = database
}
