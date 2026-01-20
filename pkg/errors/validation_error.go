package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(data interface{}) []ValidationError {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			Field:   err.Field(),
			Message: fmt.Sprintf("field '%s' failed validation: %s", err.Field(), err.Tag()),
		})
	}

	return validationErrors
}
