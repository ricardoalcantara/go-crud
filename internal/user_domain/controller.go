package userdomain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api")
	routes.POST("/user", h.Empty)
	routes.GET("/user/me", h.Empty)
}

/*
/api/user
/api/user/me
*/
