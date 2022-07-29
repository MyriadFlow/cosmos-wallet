package pasetomiddleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PASETO(t *testing.T) {
	appinit.Init()
	db := store.DB
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	newUser := user.User{
		WalletAddress: testWallet.WalletAddress,
	}
	err := db.Model(&user.User{}).Create(&newUser).Error
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Should return 200 with correct PASETO", func(t *testing.T) {
		token, err := paseto.GetPasetoForUser(testWallet.WalletAddress)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode, "status code should be 200 (OK), body: %s", rr.Body)
	})

	t.Run("Should return 401 with incorret PASETO", func(t *testing.T) {
		os.Setenv("PASETO_PRIVATE_KEY", "some invalid token")
		token, err := paseto.GetPasetoForUser(testWallet.WalletAddress)
		if err != nil {
			t.Fatal(err)
		}
		os.Setenv("PASETO_PRIVATE_KEY", "other token as valid")
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode, "status code should be 401 (Unauthorized), body: %s", rr.Body)
	})

	t.Run("Should return 401 for deleted user", func(t *testing.T) {
		// Non existance wallet addr in database
		nonExistanceWallet := testingcommon.GenerateWallet()
		token, err := paseto.GetPasetoForUser(nonExistanceWallet.WalletAddress)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode, "status code should be 401 (Unauthorized), body: %s", rr.Body)
	})

	t.Run("Should return 401 and 4011 with expired PASETO", func(t *testing.T) {
		token, err := testingcommon.GetPasetoForTestUser(testWallet.WalletAddress, 2*time.Second)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 4)

		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode, "status code should be 401 (Unauthorized), body: %s", rr.Body)
		var response httpo.ApiResponse
		body := rr.Body
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, httpo.TokenExpired, response.StatusCode, "api response status code should be 4031 (TokenExpired), body: %s", rr.Body)
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
	ginTestApp.Use(PASETO)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
