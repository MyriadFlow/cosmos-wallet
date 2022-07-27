package httpo

import (
	"github.com/gin-gonic/gin"
)

//TODO custom status codes
func SendResponse(c *gin.Context, statusCode int, apiRes ApiResponse) {
	c.JSON(statusCode, apiRes)
}
