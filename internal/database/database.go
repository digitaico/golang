package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new DB connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString = fmt.Sprintf("host=%s port=%s dbname=%s password=%s ssl")

	return nil, nil
}
