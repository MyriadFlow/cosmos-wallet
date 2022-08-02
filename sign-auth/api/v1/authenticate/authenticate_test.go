package authenticate

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/helpers/httpo"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/api/v1/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	arbitraryverify "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/cosmos_blockchain/arbitrary_verify"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/testingcommon"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostAuthenticate(t *testing.T) {
	appinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())

	url := "/api/v1.0/authenticate"
	t.Run("Should return 200 with correct wallet address", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()
		eula, flowId := callFlowIdApi(testWallet.WalletAddress, t)
		signature := getSignature(eula, flowId, testWallet)
		pubKeyBytes := testWallet.PrivKey.PubKey().Bytes()
		pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKeyBytes)
		body := AuthenticateRequest{Signature: signature, FlowId: flowId, PublicKey: pubKeyBase64}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		//Request with signature created from correct wallet address
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}

		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		authenticate(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String(), "status code should be 200 (OK), body: %s", rr.Body)
	})
	t.Run("Should return 401 with different wallet address", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()
		eula, flowId := callFlowIdApi(testWallet.WalletAddress, t)
		// Different private key will result in different wallet address
		differentWallet := testingcommon.GenerateWallet()
		signature := getSignature(eula, flowId, differentWallet)
		body := AuthenticateRequest{Signature: signature, FlowId: flowId, PublicKey: base64.StdEncoding.EncodeToString(testWallet.PrivKey.PubKey().Bytes())}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		//Request with signature stil created from different walletAddress
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		authenticate(c)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode, rr.Body.String(), "status code should be 401 (Unauthorized), body: %s", rr.Body)
	})

	t.Run("Should return 401 if flow ID doesn't exist", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()
		eula, _ := callFlowIdApi(testWallet.WalletAddress, t)

		// Non existance flow id
		randomFlowId := uuid.NewString()
		signature := getSignature(eula, randomFlowId, testWallet)
		pubKeyBytes := testWallet.PrivKey.PubKey().Bytes()
		pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKeyBytes)
		body := AuthenticateRequest{Signature: signature, FlowId: randomFlowId, PublicKey: pubKeyBase64}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		//Request with signature created from correct wallet address
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}

		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		authenticate(c)
		assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode, rr.Body.String(), "status code should be 404 (NotFound), body: %s", rr.Body)
	})

}

func callFlowIdApi(walletAddress string, t *testing.T) (eula string, flowidString string) {
	// Call flowid api
	u, err := url.Parse("/api/v1.0/flowid")
	q := url.Values{}
	logo.Infof("walletAddress: %v\n", walletAddress)
	q.Set("walletAddress", walletAddress)
	u.RawQuery = q.Encode()
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", u.String(), nil)
	req.URL.RawQuery = q.Encode()
	if err != nil {
		t.Error(err)
	}
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	flowid.GetFlowId(c)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode, "status code should be 200 (OK), body: %s", rr.Body)
	var flowIdPayload flowid.GetFlowIdPayload
	var res httpo.ApiResponse
	decoder := json.NewDecoder(rr.Result().Body)
	err = decoder.Decode(&res)
	testingcommon.ExtractPayload(&res, &flowIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	return flowIdPayload.Eula, flowIdPayload.FlowId
}

func getSignature(eula string, flowId string, testWallet testingcommon.TestWallet) string {
	messageData := eula + flowId
	messageComposed, err := arbitraryverify.ComposeArbitraryMsg(testWallet.WalletAddress, messageData)
	if err != nil {
		log.Fatal(err)
	}
	signatureBytes, err := testWallet.PrivKey.Sign([]byte(messageComposed))
	if err != nil {
		log.Fatal(err)
	}
	signature := base64.StdEncoding.EncodeToString(signatureBytes)
	return signature
}
