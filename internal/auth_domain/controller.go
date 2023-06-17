package authdomain

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
	routes.POST("/auth/token", h.Empty)
	routes.POST("/auth/refresh_token", h.Empty)
	routes.POST("/auth/logout", h.Empty)
	routes.POST("/auth/password_recovery", h.Empty)
	routes.PUT("/auth/password_recovery", h.Empty)
	routes.GET("/auth/claims", h.Empty)
}

/*
/api/user
/api/user/me
*/

/*
/api/auth/token
/api/auth/refresh_token
/api/auth/logout
/api/auth/password_recovery
/api/auth/claims
*/
