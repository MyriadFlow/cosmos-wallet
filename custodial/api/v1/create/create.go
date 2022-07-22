package create

import (
	"encoding/base64"

	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
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
		httpo.NewInternalServerError(c, "failed to create user")
		return
	}

	pubKeyBase64 := base64.StdEncoding.EncodeToString((*pubKey).Bytes())
	payload := CreatePayload{
		UserId:    userId,
		PublicKey: pubKeyBase64,
	}
	httpo.SuccessResponse(c, "Token generated successfully", payload)
}
