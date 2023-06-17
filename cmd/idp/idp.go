package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	authdomain "github.com/ricardoalcantara/go-crud/internal/auth_domain"
	"github.com/ricardoalcantara/go-crud/internal/models"
	"github.com/ricardoalcantara/go-crud/internal/setup"
	userdomain "github.com/ricardoalcantara/go-crud/internal/user_domain"
)

func init() {
	setup.Env()
}

func main() {

	log.Info("Server Started")

	db, err := models.Database()

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	name := "Ricardo"
	email := "alcantarafox@yahoo.com.br"
	username := "ricardo"
	password := hashAndSalt([]byte("ricardo"))
	user := models.User{Name: &name, Email: &email, Username: &username, Password: &password}
	_ = db.Create(&user)

	var q_user models.User
	db.First(&q_user)

	log.WithFields(log.Fields{
		"Id":       q_user.ID,
		"Name":     *q_user.Name,
		"Password": *q_user.Password,
	}).Info("First")

	r := gin.Default()

	authdomain.RegisterRoutes(r, db)
	userdomain.RegisterRoutes(r, db)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run()
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
