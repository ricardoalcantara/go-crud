package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/ricardoalcantara/go-crud/internal/models"
	"github.com/ricardoalcantara/go-crud/internal/setup"
)

func init() {
	setup.Env()
}

func main() {
	log.Info("Migrate Started")

	db, err := models.Database()

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&models.UserPasswordRecovery{}); err != nil {
		log.Fatal(err)
	}

	log.Info("Migrate Finished")
}
