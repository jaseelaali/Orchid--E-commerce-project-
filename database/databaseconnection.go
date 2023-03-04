package database

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() {
	// load database
	if err := godotenv.Load(); err != nil {
		fmt.Println("error in loading env file")
	}
	// connect database
	var DB *gorm.DB
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error in connecting database")
	}
	log.Println("Successfully connected to database")
	//sync database
	err = DB.AutoMigrate(
		&models.Admin{},
		&models.User{},
	)
	if err != nil {
		log.Println("error in syncing the database")
	}

}