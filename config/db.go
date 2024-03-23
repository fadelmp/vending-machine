package config

import (
	"fmt"
	"os"
	"vending-machine/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// dbSchema := os.Getenv("DB_SCHEMA")

	// Create the connection string
	// connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s search_path=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass, dbSchema)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		logrus.Error("Cannot Connect to PostgreSQL DB")
		panic(err)
	}

	createTable(db)
	migrateDDL(db)

	return db
}

func createTable(db *gorm.DB) {
	db.CreateTable(&domain.Vending{})
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&domain.Vending{})
}
