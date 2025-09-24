package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdinnesh/prepaid-card/services/auth/config"
	"gorm.io/gorm"
)

func New(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// test route
	router.GET("/auth-test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Auth Service is running",
		})
	})

	router.Group("/api/v1")
	return router
}
