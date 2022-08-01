// Package create provides Api methods to create user wallet and store the mnemonic in database
package create

import (
	"encoding/base64"
	"net/http"

	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	walletaddress "github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos/wallet_address"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies /authenticate to gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/create")
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

	userWalletAddr, err := walletaddress.GetWalletAddrFromPubKey(env.MustGetEnv("WALLET_ADDRESS_HRP"), *pubKey)
	if err != nil {
		logo.Errorf("failed to get wallet address from public key of user with id %s: %s", userId, err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get user walletaddress").
			Send(c, http.StatusInternalServerError)
		return
	}
	// Convert the public key to base64 to send it as JSON
	pubKeyBase64 := base64.StdEncoding.EncodeToString((*pubKey).Bytes())
	payload := CreatePayload{
		UserId:     userId,
		PublicKey:  pubKeyBase64,
		WalletAddr: userWalletAddr,
	}

	httpo.NewSuccessResponse(http.StatusOK, "User created successfully", payload).
		Send(c, http.StatusOK)
}
