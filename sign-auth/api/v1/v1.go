package apiv1

import (
	pasetomiddleware "github.com/MyriadFlow/cosmos-wallet/sign-auth/api/middleware/auth/paseto"
	authenticate "github.com/MyriadFlow/cosmos-wallet/sign-auth/api/v1/authenticate"
	flowid "github.com/MyriadFlow/cosmos-wallet/sign-auth/api/v1/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/api/v1/healthcheck"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /v1.0 group and all child routes to given gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)
		healthcheck.ApplyRoutes(v1)
		v1.Use(pasetomiddleware.PASETO)
	}
}
