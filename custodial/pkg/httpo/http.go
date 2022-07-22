package httpo

import (
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

func ErrResponse(c *gin.Context, statusCode int, errMessage string) {
	response := ApiResponse{
		StatusCode: statusCode,
		Error:      errMessage,
	}
	c.JSON(response.StatusCode, response)
}

//TODO: document methods
func CErrResponse(c *gin.Context, statusCode int, customStatusCode int, errMessage string) {
	response := ApiResponse{
		StatusCode: customStatusCode,
		Error:      errMessage,
	}
	c.JSON(statusCode, response)
}

func SuccessResponse(c *gin.Context, message string, payload interface{}) {
	response := ApiResponse{
		StatusCode: http.StatusOK,
		Payload:    payload,
		Message:    message,
	}
	c.JSON(response.StatusCode, response)
}

func InternalServerError(c *gin.Context) {
	response := ApiResponse{
		StatusCode: http.StatusInternalServerError,
		Error:      "unexpected error occurred",
	}
	c.JSON(response.StatusCode, response)
}

func NewInternalServerError(c *gin.Context, format string, args ...interface{}) {
	logo.Errorf(format, args...)
	response := ApiResponse{
		StatusCode: http.StatusInternalServerError,
		Error:      "unexpected error occurred",
	}
	c.JSON(response.StatusCode, response)
}
