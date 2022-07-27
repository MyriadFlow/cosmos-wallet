package authenticate

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	flowidmethods "github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid/flowid_methods"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {
	var req AuthenticateRequest
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to validate body").
			Send(c, http.StatusBadRequest)
		return
	}

	bytesPubKey, err := base64.StdEncoding.DecodeString(req.PublicKey)
	if err != nil {
		logo.Errorf("failed to decode base64 public key: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to decode base64 public key").
			Send(c, http.StatusBadRequest)
		return
	}

	pubKey := secp256k1.PubKey{
		Key: bytesPubKey,
	}
	pasetoToken, err := flowidmethods.VerifySignAndGetPaseto(pubKey, req.Signature, req.FlowId)
	if err != nil {
		logo.Errorf("failed to get paseto: %s", err)

		if errors.Is(err, flowidmethods.ErrSignDenied) {
			httpo.NewErrorResponse(http.StatusUnauthorized, "signature denied").
				Send(c, http.StatusUnauthorized)
		}
		httpo.NewErrorResponse(500, "failed to verify and get paseto").Send(c, 500)
		return
	} else {
		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		httpo.NewSuccessResponse(http.StatusOK, "Token generated successfully", payload).Send(c, http.StatusOK)
	}
}
