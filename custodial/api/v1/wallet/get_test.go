package wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	usermethods "github.com/MyriadFlow/cosmos-wallet/custodial/models/user/user_methods"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostAuthenticate(t *testing.T) {
	appinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())

	url := "/api/v1.0/wallet"
	t.Run("Should return 200 and public key", func(t *testing.T) {
		rr := httptest.NewRecorder()
		_, userId, err := usermethods.Create()
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		body := GetWalletRequest{UserId: userId}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		//Request with signature created from correct wallet address
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = req
		getWallet(c)
		assert.Equal(t, http.StatusOK, rr.Code, "status code should be 200 (OK), body: %s", rr.Body)
	})

}
