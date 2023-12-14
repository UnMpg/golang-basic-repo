package handler

import (
	"go-project/api/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Uusecase usecase.UserUsecase
}

func NewUserHandler(userRoute *gin.RouterGroup, userUc usecase.UserUsecase) {
	handle := UserHandler{Uusecase: userUc}

	userRoute.POST("/register", handle.RegisterUser)
	userRoute.GET("/tes", handle.CobaTest)
}
