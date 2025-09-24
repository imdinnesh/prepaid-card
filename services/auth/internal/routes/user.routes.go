package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imdinnesh/prepaid-card/services/auth/config"
	"github.com/imdinnesh/prepaid-card/services/auth/internal/handler"
	"github.com/imdinnesh/prepaid-card/services/auth/internal/repository"
	"github.com/imdinnesh/prepaid-card/services/auth/internal/services"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	auth := r.Group("/auth")
	auth.POST("/register", userHandler.RegisterUser)
}
