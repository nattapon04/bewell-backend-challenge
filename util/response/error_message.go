package response

import "net/http"

const (
	ErrValidationFailed = `VALIDATION_FAILED`
	ErrInvalidInput     = `INVALID_INPUT`
	ErrNotFound         = `NOT_FOUND`
	ErrSomethingWrong   = `SOMETHING_WRONG`
	ErrBadRequest       = `BAD_REQUEST`
)

var ErrorStatusText = map[string]string{
	ErrValidationFailed: `Validation Failed`,
	ErrNotFound:         `Not Found`,
	ErrSomethingWrong:   `Something Wrong`,
}

var ErrorMessage = map[string]string{
	ErrInvalidInput:   `Invalid format of the input`,
	ErrNotFound:       `Data Not Found`,
	ErrSomethingWrong: `Something Wrong`,
}

var ErrorsStatusCode = map[string]int{
	ErrSomethingWrong:   http.StatusInternalServerError,
	ErrValidationFailed: http.StatusUnprocessableEntity,
	ErrBadRequest:       http.StatusBadRequest,
	ErrNotFound:         http.StatusNotFound,
}
