package handler

import (
	"go-project/models"
	"go-project/utils/log"
	"go-project/utils/message"
	"go-project/utils/validator"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Success 201 {object} models.User
// @Router /user/api/register [post]
// @Produce json
func (Uuc *UserHandler) RegisterUser(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "RegisterUser", "request")
	defer span.End()
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

	data, err := Uuc.Uusecase.CreateUser(ctx, req)
	if err != nil {
		c.JSON(message.StatusBadRequestCode, models.CreateResponse(message.StatusBadRequestCode, message.FAILED, "Failed Create User", nil))
		return
	}

	c.JSON(message.StatusOk, models.CreateResponse(message.StatusOk, message.SUCCESS, message.SUCCESS, data))

}

func (Uuc *UserHandler) CobaTest(c *gin.Context) {
	data, err := Uuc.Uusecase.CobaUsecaseTesting()
	if err != nil {
		return
	}
	c.JSON(message.StatusOk, models.CreateResponse(message.StatusOk, message.SUCCESS, message.SUCCESS, data))
}
