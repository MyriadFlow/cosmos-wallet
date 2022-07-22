package tokenauthmiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_TOKENAUTH(t *testing.T) {
	appinit.Init()
	//TODO package to set env if not set for testing
	gin.SetMode(gin.TestMode)

	t.Run("Should return 200 with correct TOKEN", func(t *testing.T) {
		token := env.MustGetEnv("AUTH_TOKEN")
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	t.Run("Should return 401 with incorret TOKEN", func(t *testing.T) {
		token := "some invalid token"
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	})
}

func callApi(t *testing.T, token string) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ginTestApp := gin.New()

	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", token)
	ginTestApp.Use(TOKENAUTH)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
