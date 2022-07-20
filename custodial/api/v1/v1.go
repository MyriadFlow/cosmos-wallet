package apiv1

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/create"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/healthcheck"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		create.ApplyRoutes(v1)
		healthcheck.ApplyRoutes(v1)
	}
}
