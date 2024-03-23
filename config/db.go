package config

import (
	"fmt"
	"os"
	"vending-machine/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", connStr)
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
