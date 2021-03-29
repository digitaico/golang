package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new DB")
	return nil, nil
}
