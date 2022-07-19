package flowid

import (
	"net/http"
	"strings"

	usermethods "github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/httpo"
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
		httpo.ErrResponse(c, http.StatusBadRequest, "wallet address (walletAddress) is required")
		return
	}

	_, _, err := bech32.DecodeAndConvert(walletAddress)
	if err != nil {
		log.Errorf("failed to decode bech32 wallet address %s: %s", walletAddress, err)
		httpo.ErrResponse(c, http.StatusBadRequest, "failed to parse bech32 Wallet address (walletAddress)")
		return
	}

	flowId, err := usermethods.CreateFlowId(walletAddress)
	if err != nil {
		log.Errorf("failed to generate flow id: %s", err)
		httpo.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	userAuthEULA := env.MustGetEnv("AUTH_EULA")
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.SuccessResponse(c, "Flowid successfully generated", payload)
}
