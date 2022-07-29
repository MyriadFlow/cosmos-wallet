package httpo

import "github.com/gin-gonic/gin"

// ApiResponse defines struct used for http response
type ApiResponse struct {
	// Custom status code
	StatusCode int `json:"status,omitempty"`

	Error string `json:"error,omitempty"`

	// Message in case of success
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// Sends ApiResponse with gin context and standard statusCode
func (apiRes *ApiResponse) Send(c *gin.Context, statusCode int) {
	c.JSON(statusCode, apiRes)
}

// NewSuccessResponse returns ApiResponse for success
func NewSuccessResponse(customStatusCode int, message string, payload interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: customStatusCode,
		Message:    message,
		Payload:    payload,
	}
}

// NewSuccessResponse returns ApiResponse for error
func NewErrorResponse(customStatusCode int, errorStr string) *ApiResponse {
	return &ApiResponse{
		StatusCode: customStatusCode,
		Error:      errorStr,
	}
}
