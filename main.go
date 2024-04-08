package myvalidates

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// will validate all struct variabel whose has gorm:"validate"
func ValidateReq(request interface{}) error {
	var errString string

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
// ValidateCustom(request, "Mama") -> use like this
// ValidateCustom(request, "Mama", "Papa") -> and can like this
func ValidateCustom(request interface{}, fieldToValidate ...string) error {
	var errString string

	err := validate.StructPartial(request, fieldToValidate...)
	if err != nil {
		es := err.(validator.ValidationErrors)
		for _, errVali := range es {
			errString = errVali.Field() + " is " + errVali.Tag()
			break
		}
		return errors.New(errString)
	}

	return nil
}

// this is the opposite ValidateCustom
func ValidateCustomExcept(request interface{}, fieldsToExceptValidate ...string) error {
	var errString string
	err := validate.StructExcept(request, fieldsToExceptValidate...)
	if err != nil {
		es := err.(validator.ValidationErrors)
		for _, errVali := range es {
			errString = errVali.Field() + " is " + errVali.Tag()
			break
		}
		return errors.New(errString)
	}

	return nil
}
