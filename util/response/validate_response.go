package response

import (
	"bewell-backend-challenge/internal/model"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type fieldError struct {
	err validator.FieldError
}

func parseValidateMessage(err error) *model.ValidateError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) { //nolint:nolintlint,forcetypeassert,errorlint
			validationErr := &model.ValidateError{
				ErrorCode: strings.ToUpper(err.ActualTag()),
				Field:     err.Field(),
				Message:   fieldError{err}.String(),
			}

			return validationErr
		}
	} else {
		validationErr := &model.ValidateError{
			ErrorCode: ErrInvalidInput,
			Message:   err.Error(),
		}

		return validationErr
	}

	return nil
}

func (q fieldError) String() string {
	var stringsBuilder strings.Builder

	stringsBuilder.WriteString("Validation failed on field '" + q.err.Field() + "'")
	stringsBuilder.WriteString(", condition: " + q.err.ActualTag())

	if q.err.Param() != "" {
		stringsBuilder.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		stringsBuilder.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return stringsBuilder.String()
}
