package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var DB *gorm.DB

func Db_details() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
}

func InitDatabase() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Something went wrong while reading .env file")
	}

	dsn := Db_details()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&User{})

	if err != nil {
		log.Fatal("Something went wrong while connecting to the database")
	}

	fmt.Println("Connected to the database")
}
