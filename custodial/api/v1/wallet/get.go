package getwallet

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
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
		httpo.ErrResponse(c, http.StatusBadRequest, "failed to fetch user")
		return
	}
	userWallet, err := user.Get(req.UserId)

	if err != nil {
		if errors.Is(err, errorso.ErrRecordNotFound) {
			httpo.ErrResponse(c, http.StatusNotFound, "user with given id not found")
			return
		}
		logo.Errorf("failed to fetch user with id %s: %s", req.UserId, err)
		httpo.NewInternalServerError(c, "failed to fetch user")
		return
	}

	privKey, err := blockchain_cosmos.GetWallet(userWallet.Mnemonic)
	if err != nil {
		logo.Errorf("failed to get wallet for user with id %s: %s", req.UserId, err)
		httpo.NewInternalServerError(c, "failed to get user public key")
		return
	}
	pubKeyBase64 := base64.StdEncoding.EncodeToString(privKey.PubKey().Bytes())
	payload := GetWalletPayload{
		PublicKey: pubKeyBase64,
	}
	httpo.SuccessResponse(c, "User fetched successfully", payload)
}
