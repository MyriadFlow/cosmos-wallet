// Package create provides Api methods to create user wallet and store the mnemonic in database
package create

import (
	"encoding/base64"
	"net/http"

	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies /authenticate to gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", create)
	}
}

func create(c *gin.Context) {
	pubKey, userId, err := usermethods.Create()
	if err != nil {
		logo.Errorf("failed to create user: %s", err)
		httpo.NewErrorResponse(500, "failed to create user").
			Send(c, 500)
		return
	}

	// Convert the public key to base64 to send it as JSON
	pubKeyBase64 := base64.StdEncoding.EncodeToString((*pubKey).Bytes())
	payload := CreatePayload{
		UserId:    userId,
		PublicKey: pubKeyBase64,
	}

	httpo.NewSuccessResponse(http.StatusOK, "Token generated successfully", payload).
		Send(c, http.StatusOK)
}
