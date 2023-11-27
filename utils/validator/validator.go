package validator

import (
	"go-project/utils/log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	log.Log.Info("Init validator")

	Validator = validator.New()
}

func ReqUsernameValidate(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	// if !LenStringInput(name, 10) {
	// 	return false
	// }

	return StringRegex(name)

}

func ReqStringValidate(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	return StringRegex(name)

}

func ReqEmailValidate(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	return EmailRegex(email)

}

func ReqNumberValidate(fl validator.FieldLevel) bool {
	number := fl.Field().String()

	return NumberRegex(number)

}

func ReqStringNumberValidate(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	return StringNumberRegex(input)

}

func LenStringInput(input string, max int) bool {
	if len(input) <= max {
		return true
	}
	return true
}

func StringRegex(input string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z]+$`)
	result := regex.MatchString(input)
	return result
}

func EmailRegex(input string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	result := regex.MatchString(input)
	return result
}

func NumberRegex(input string) bool {
	regex, _ := regexp.Compile(`^[0-9]+$`)
	result := regex.MatchString(input)
	return result
}

func StringNumberRegex(input string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9.,_@/:[:space:]-]+$`)
	result := regex.MatchString(input)
	return result
}
