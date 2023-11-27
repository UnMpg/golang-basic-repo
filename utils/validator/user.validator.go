package validator

import (
	"fmt"
	"go-project/models"

	"github.com/go-playground/validator/v10"
)

func UserRegis(req models.User) models.ValidateMessage {
	Validator.RegisterValidation("ReqUsername", ReqUsernameValidate)
	Validator.RegisterValidation("ReqEmail", ReqEmailValidate)
	Validator.RegisterValidation("ReqStringNumberChar", ReqStringNumberValidate)

	var mes models.ValidateMessage
	if err := Validator.Struct(req); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			mes.Message = append(mes.Message, fmt.Sprintf("Validation failed on field '%s' with value '%s'.", err.Field(), err.Value()))
		}
	}

	return mes
}
