package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateDatabase initializes database connection
func CreateDatabase() *gorm.DB {
	dbPass := os.Getenv("DB_PASSWORD")

	fmt.Println("DB_PASSWORD: " + dbPass)

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"postgres", "5432", "restapi", "restapi", dbPass, "disable")

	fmt.Println(connectionString)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	db.LogMode(true)

	db = db.Set("gorm:auto_preload", true)

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})

	return db
}
