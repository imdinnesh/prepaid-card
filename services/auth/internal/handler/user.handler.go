package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdinnesh/prepaid-card/services/auth/internal/services"
)

type UserHandler interface{
	RegisterUser()
}

type userHandler struct{
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler{
	return &userHandler{
		userService:userService,
	}		
}

func (h *userHandler) RegisterUser(ctx *gin.Context){
	var req struct{
		Phone string `json:"phone" binding:"required"`
	}

	if err:=ctx.ShouldBindJSON(&req);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	err:=h.userService.RegisterUser(req.Phone)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"User registered successfully",
	})	

}
	





