package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	db_url := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
