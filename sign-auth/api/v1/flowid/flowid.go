package flowid

import (
	"net/http"
	"strings"

	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	usermethods "github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.GET("", GetFlowId)
	}
}

func GetFlowId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" || !strings.HasPrefix(walletAddress, env.MustGetEnv("WALLET_ADDRESS_HRP")) {
		httpo.NewErrorResponse(http.StatusBadRequest, "wallet address (walletAddress) is required or is not valid").
			Send(c, http.StatusBadRequest)
		return
	}

	_, _, err := bech32.DecodeAndConvert(walletAddress)
	if err != nil {
		log.Errorf("failed to decode bech32 wallet address %s: %s", walletAddress, err)
		httpo.NewErrorResponse(httpo.WalletAddressInvalid, "failed to parse bech32 Wallet address (walletAddress)").Send(c, http.StatusBadRequest)
		return
	}

	flowId, err := usermethods.CreateFlowId(walletAddress)
	if err != nil {
		log.Errorf("failed to generate flow id: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").Send(c, http.StatusInternalServerError)
		return
	}
	userAuthEULA := env.MustGetEnv("AUTH_EULA")
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.NewSuccessResponse(http.StatusOK, "Flowid successfully generated", payload).Send(c, http.StatusOK)
}
