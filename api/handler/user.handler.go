package handler

import (
	"go-project/models"
	"go-project/utils/log"
	"go-project/utils/message"
	"go-project/utils/validator"

	"github.com/gin-gonic/gin"
)

func (Uuc *UserHandler) RegisterUser(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Log.Error(log.Register, "Error Binding Json req ", err)
		c.JSON(message.StatusBadRequestCode, models.CreateResponse(message.StatusBadRequestCode, message.FAILED, "Binding Failed", nil))
		return
	}

	validateReq := validator.UserRegis(req)

	if validateReq.Message != nil {
		c.JSON(message.StatusBadRequestCode, models.CreateResponse(message.StatusBadRequestCode, message.FAILED, "Error validate Input", validateReq))
		return
	}

	data, err := Uuc.Uusecase.CreateUser(req)
	if err != nil {
		c.JSON(message.StatusBadRequestCode, models.CreateResponse(message.StatusBadRequestCode, message.FAILED, "Failed Create User", nil))
		return
	}

	c.JSON(message.StatusOk, models.CreateResponse(message.StatusOk, message.SUCCESS, message.SUCCESS, data))

}
