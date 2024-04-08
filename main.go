package myvalidates

import (
	"errors"

	"github.com/go-playground/validator"
)

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

func ValidateCustom(request interface{}, arr []string) error {
	var errString string
	validate := validator.New()

	for _, key := range arr {
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
