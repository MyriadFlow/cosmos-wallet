package authenticate

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	flowidmethods "github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid/flowid_methods"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/httpo"
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
		httpo.ErrResponse(c, http.StatusBadRequest, "failed to validate body")
		return
	}

	bytesPubKey, err := base64.StdEncoding.DecodeString(req.PublicKey)
	if err != nil {
		logo.Errorf("failed to decode base64 public key: %s", err)
		httpo.ErrResponse(c, http.StatusBadRequest, "failed to decode base64 public key")
		return
	}

	pubKey := secp256k1.PubKey{
		Key: bytesPubKey,
	}
	pasetoToken, err := flowidmethods.VerifySignAndGetPaseto(pubKey, req.Signature, req.FlowId)
	if err != nil {
		logo.Errorf("failed to get paseto: %s", err)

		if errors.Is(err, flowidmethods.ErrSignDenied) {
			httpo.ErrResponse(c, http.StatusUnauthorized, "signature denied")
		}
		httpo.ErrResponse(c, 500, "failed to verify and get paseto")
		return
	} else {
		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		httpo.SuccessResponse(c, "Token generated successfully", payload)
	}
}
