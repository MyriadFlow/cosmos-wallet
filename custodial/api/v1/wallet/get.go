// Package create provides Api methods to get user wallet's public key
package getwallet

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/wallet")
	{
		g.POST("", getWallet)
	}
}

func getWallet(c *gin.Context) {
	var req GetWalletRequest
	err := c.BindJSON(&req)
	if err != nil {
		logo.Errorf("failed to bind json: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "request body is invalid").
			Send(c, http.StatusBadRequest)
		return
	}
	userWallet, err := user.Get(req.UserId)

	if err != nil {
		// If user doesn't exist
		if errors.Is(err, errorso.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user with given id not found").
				Send(c, http.StatusNotFound)
			return
		}

		// Unexpected error has occured
		logo.Errorf("failed to fetch user with id %s: %s", req.UserId, err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").
			Send(c, http.StatusInternalServerError)
		return
	}

	// Get private key from mnemonic
	privKey, err := blockchain_cosmos.GetPrivKey(userWallet.Mnemonic)
	if err != nil {
		logo.Errorf("failed to get wallet for user with id %s: %s", req.UserId, err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get user public key").
			Send(c, http.StatusInternalServerError)
		return
	}

	// Convert the public key to base64 to send it as JSON
	pubKeyBase64 := base64.StdEncoding.EncodeToString(privKey.PubKey().Bytes())
	payload := GetWalletPayload{
		PublicKey: pubKeyBase64,
	}

	httpo.NewSuccessResponse(http.StatusOK, "User fetched successfully", payload).
		Send(c, http.StatusOK)
}
