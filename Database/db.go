package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func Db_details() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
}

func InitDatabase() error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	dsn := Db_details()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		return err
	}

	return nil
}
