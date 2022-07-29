// Package healthcheck provides healthcheck Api to check if the app is alive,
// returns 200 if it is not in crashed state
package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/healthcheck")
	{
		g.GET("", healthCheck)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "alive",
	})
}
