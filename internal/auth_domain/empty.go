package authdomain

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Empty(c *gin.Context) {
	// h.DB
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
