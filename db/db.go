package db

import (
	"fmt"
	"log"
	"os"

	"github.com/LipMark/petProject1ADB/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func getEnv(key, identify string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return identify
}

func Init() {
	user := getEnv("PG_USER", "user123")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "8080")
	database := getEnv("PG_DB", "tasks")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.Task{}) {
		err := db.CreateTable(&models.Task{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Task{})
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
