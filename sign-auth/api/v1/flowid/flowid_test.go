package flowid

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetFlowId(t *testing.T) {
	appinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	u, err := url.Parse("/api/v1.0/flowid")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Should fail if wallet address is not valid", func(t *testing.T) {
		t.Run("wallet address doesn't have cosmos prefix", func(t *testing.T) {
			q := url.Values{}
			q.Set("walletAddress", "abc")
			u.RawQuery = q.Encode()
			rr := httptest.NewRecorder()

			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				t.Error(err)
			}
			c, _ := gin.CreateTestContext(rr)
			c.Request = req
			GetFlowId(c)
			assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode, "status code should be 400 (BadRequest), body: %s", rr.Body)
		})

		t.Run("wallet address is not valid bech32", func(t *testing.T) {
			q := url.Values{}
			q.Set("walletAddress", "cosmosabc")
			u.RawQuery = q.Encode()
			rr := httptest.NewRecorder()

			req, err := http.NewRequest("GET", u.String(), nil)
			if err != nil {
				t.Error(err)
			}
			c, _ := gin.CreateTestContext(rr)
			c.Request = req
			GetFlowId(c)
			assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode, "status code should be 400 (BadRequest), body: %s", rr.Body)
		})

	})

	t.Run("Should be able to get flow id", func(t *testing.T) {

		q := url.Values{}
		q.Set("walletAddress", testWalletAddress)
		u.RawQuery = q.Encode()
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			t.Error(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		GetFlowId(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode, "status code should be 200 (OK), body: %s", rr.Body)
	})

}
