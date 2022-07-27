package apiv1

import (
	tokenauthmiddleware "github.com/MyriadFlow/cosmos-wallet/custodial/api/middleware/auth/tokenauth"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/create"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/healthcheck"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/transfer"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		healthcheck.ApplyRoutes(v1)
		v1.Use(tokenauthmiddleware.TOKENAUTH)
		create.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
	}
}
