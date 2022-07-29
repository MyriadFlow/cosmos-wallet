// Package transfer provides Api methods to transfer tokens from the wallet to another wallet on same chain
package transfer

import (
	"net/http"

	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin RouterGroup
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/transfer")
	{
		g.POST("", transfer)
	}
}

func transfer(c *gin.Context) {
	var req TransferRequest
	err := c.BindJSON(&req)
	if err != nil {
		logo.Errorf("failed to bind json: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "request body is not valid").
			Send(c, http.StatusBadRequest)
		return
	}

	txHash, err := usermethods.Transfer(req.UserId, req.From, req.To, req.Amount)
	if err != nil {
		logo.Errorf("failed to transfer tokens for user with id %s: %s", req.UserId, err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to transfer tokens").
			Send(c, http.StatusInternalServerError)
		return
	}

	payload := TransferPayload{
		TransactionHash: txHash,
	}
	httpo.NewSuccessResponse(http.StatusOK, "Transfer transaction broadcasted", payload).
		Send(c, http.StatusOK)
}
