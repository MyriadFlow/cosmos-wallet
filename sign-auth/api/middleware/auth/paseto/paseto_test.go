package pasetomiddleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	customstatuscodes "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/constants/custom_status_codes"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/httpo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PASETO(t *testing.T) {
	appinit.Init()
	//TODO package to set env if not set for testing
	db := store.DB
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := "cosmos1v7lktr9l6vx7sqz5wc9t0v0dq7gsn2zglkn9vx"
	newUser := user.User{
		WalletAddress: testWalletAddress,
	}
	err := db.Model(&user.User{}).Create(&newUser).Error
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Should return 200 with correct PASETO", func(t *testing.T) {
		token, err := paseto.GetPasetoForUser(testWalletAddress)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	t.Run("Should return 401 with incorret PASETO", func(t *testing.T) {
		os.Setenv("PASETO_PRIVATE_KEY", "some invalid token")
		token, err := paseto.GetPasetoForUser(testWalletAddress)
		if err != nil {
			t.Fatal(err)
		}
		os.Setenv("PASETO_PRIVATE_KEY", "other token as valid")
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	})

	t.Run("Should return 401 and 4011 with expired PASETO", func(t *testing.T) {
		token, err := testingcommon.GetPasetoForTestUser(testWalletAddress, 2*time.Second)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 4)

		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
		var response httpo.ApiResponse
		body := rr.Body
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, customstatuscodes.TokenExpired, response.StatusCode)
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