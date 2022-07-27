package transfer

import (
	"net/http"

	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
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
		httpo.ErrResponse(c, http.StatusBadRequest, "failed to fetch user")
		return
	}
	txHash, err := usermethods.Transfer(req.UserId, req.From, req.To, req.Amount)
	if err != nil {
		logo.Errorf("failed to transfer tokens for user with id %s: %s", req.UserId, err)
		httpo.NewInternalServerError(c, "failed to transfer tokens")
		return
	}

	payload := TransferPayload{
		TransactionHash: txHash,
	}
	httpo.SuccessResponse(c, "Transfer transaction broadcasted", payload)
}
