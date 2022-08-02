// Package api provide support to create /api group
package api

import (
	v1 "github.com/MyriadFlow/cosmos-wallet/sign-auth/api/v1"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /api group and v1 routes to given gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
