package tokenauthmiddleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
	ErrInvalidToken      = errors.New("token is not valid")
)

func TOKENAUTH(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		err = fmt.Errorf("failed to bind header, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(headers.Authorization, ErrAuthHeaderMissing)
		httpo.ErrResponse(c, http.StatusBadRequest, ErrAuthHeaderMissing.Error())
		c.Abort()
		return
	}

	if headers.Authorization != env.MustGetEnv("AUTH_TOKEN") {
		logValidationFailed(headers.Authorization, ErrInvalidToken)
		httpo.ErrResponse(c, http.StatusUnauthorized, ErrInvalidToken.Error())
		c.Abort()
		return
	}
}

func logValidationFailed(token string, err error) {
	logo.Warnf("validation failed with token %v and error: %v", token, err)
}
