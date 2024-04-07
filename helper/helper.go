package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta
	Data interface{}
}

type meta struct {
	Message string
	Code    int
	Status  string
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	actualResponse := response{
		Meta: meta,
		Data: data,
	}
	return actualResponse
}

func ApiErrorResponse(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
