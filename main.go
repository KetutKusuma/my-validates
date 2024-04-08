package myvalidates

import (
	"errors"

	"github.com/go-playground/validator"
)

// will validate all struct variabel whose has gorm:"validate"
func ValidateReq(request interface{}) error {
	var errString string
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		for _, errValidation := range err.(validator.ValidationErrors) {
			errString = errValidation.Field() + " is " + errValidation.Tag()
			break
		}
		return errors.New(errString)
	}

	return nil

}

// will validate custom from array to validate (string)
// the string is follow what name variable from your struct
// ex :
//
//	type Mama struct{
//			Mama string "json:"mama" validate:required"
//			Papa string "json:"papa" validate:required"
//	}
//
// use "Mama" not the json
// ValidateCustom(request, []string{"Mama"}) -> use like this
func ValidateCustom(request interface{}, arrToValidate []string) error {
	var errString string
	validate := validator.New()

	for _, key := range arrToValidate {
		err := validate.StructPartial(request, key)
		if err != nil {
			es := err.(validator.ValidationErrors)
			for _, errVali := range es {
				errString = errVali.Field() + " is " + errVali.Tag()
				break
			}
			return errors.New(errString)
		}
	}

	return nil
}

func ValidateCustomExcept(request interface{}, arrToExceptValidate []string) error {
	var errString string
	validate := validator.New()

	for _, key := range arrToExceptValidate {
		err := validate.StructExcept(request, key)
		if err != nil {
			es := err.(validator.ValidationErrors)
			for _, errVali := range es {
				errString = errVali.Field() + " is " + errVali.Tag()
				break
			}
			return errors.New(errString)
		}
	}

	return nil
}
