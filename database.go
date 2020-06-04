package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// CreateDatabase initializes database connection
func CreateDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// db.LogMode(true)

	db = db.Set("gorm:auto_preload", true)

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})

	return db
}
