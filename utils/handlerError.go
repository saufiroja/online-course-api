package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandlerErrorValidator(err error) error {
	var message string
	var code int

	if obj, ok := err.(validator.ValidationErrors); ok {
		for _, v := range obj {
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", v.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid", v.Field())
			case "min":
				message = fmt.Sprintf("%s is too short", v.Field())
			case "max":
				message = fmt.Sprintf("%s is too long", v.Field())

			}
		}
	}

	return HandlerError(code, message)
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Error) Error() string {
	return r.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}

func GetStatusCode(err error) int {
	if err == nil {
		return 200
	}

	if e, ok := err.(*Error); ok {
		return e.StatusCode()
	}

	return 500
}

func HandlerError(code int, message string) error {
	return &Error{
		Message: message,
		Code:    code,
	}
}
