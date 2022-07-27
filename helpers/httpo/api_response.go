package httpo

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	StatusCode int         `json:"status,omitempty"`
	Error      string      `json:"error,omitempty"`
	Message    string      `json:"message,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}

func (apiRes *ApiResponse) Send(c *gin.Context, statusCode int) {
	c.JSON(statusCode, apiRes)
}

func NewSuccessResponse(statusCode int, message string, payload interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: statusCode,
		Message:    message,
		Payload:    payload,
	}
}

func NewErrorResponse(statusCode int, errorStr string) *ApiResponse {
	return &ApiResponse{
		StatusCode: statusCode,
		Error:      errorStr,
	}
}
