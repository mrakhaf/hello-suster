package common

import (
	"net/http"
	"strconv"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()
	err := validate.RegisterValidation("int_len", validateNumberOfDigit)

	if err != nil {
		panic(err)
	}

	return &Validator{
		validator: validate,
	}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func validateNumberOfDigit(fl validator.FieldLevel) bool {
	field := fl.Field()
	param, err := strconv.Atoi(fl.Param())
	if err != nil {
		panic(err.Error())
	}

	v := field.Int()
	if v < 0 {
		panic("negative number")
	}

	n := 0
	for ; v > 0; v /= 10 {
		n += 1
	}

	return n == param
}
