package response

import (
	"bewell-backend-challenge/internal/model"
	"bewell-backend-challenge/util/helpers/appstring"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, model.ValidateMessage{
		StatusText: ErrorStatusText[ErrValidationFailed],
		Error:      parseValidateMessage(err),
	})
}

func HandleError(c *gin.Context, statusCode int, errorMessage model.Message) {
	if errorMessage.StatusText == "" {
		errorMessage.StatusText = http.StatusText(statusCode)
	}

	if errorMessage.Error != nil {
		if errorMessage.Error.ErrorCode == "" {
			errorMessage.Error.ErrorCode = strings.ToUpper(appstring.ToSnakeCase(errorMessage.StatusText))
		}
	}

	c.JSON(statusCode, errorMessage)
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}

	message := ErrorsStatusCode[err.Error()]
	if message != 0 {
		return message
	}

	return http.StatusInternalServerError
}

func GetMessage(err error) string {
	if err == nil {
		return ErrorMessage[ErrSomethingWrong]
	}

	message := ErrorMessage[err.Error()]
	if message != "" {
		return message
	}

	return err.Error()
}

func GetErrorCode(err error) string {
	if err == nil {
		return ""
	}

	if _, ok := ErrorsStatusCode[err.Error()]; ok {
		return err.Error()
	}

	if _, ok := ErrorMessage[err.Error()]; ok {
		return err.Error()
	}

	return ""
}
