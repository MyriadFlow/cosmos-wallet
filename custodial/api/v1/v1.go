// Package apiv1 provide support to create /v1.0 group and add all routes
package apiv1

import (
	tokenauthmiddleware "github.com/MyriadFlow/cosmos-wallet/custodial/api/middleware/auth/tokenauth"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/create"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/healthcheck"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/transfer"
	"github.com/MyriadFlow/cosmos-wallet/custodial/api/v1/wallet"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /v1.0 group and all child routes to given gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		healthcheck.ApplyRoutes(v1)
		v1.Use(tokenauthmiddleware.TOKENAUTH)
		create.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		wallet.ApplyRoutes(v1)
	}
}
