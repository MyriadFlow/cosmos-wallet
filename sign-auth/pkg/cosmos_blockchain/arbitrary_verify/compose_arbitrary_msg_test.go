package arbitraryverify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ComposeArbitraryMsg(t *testing.T) {
	signer := "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8"
	data := "max"
	res, err := ComposeArbitraryMsg(signer, data)
	if err != nil {
		t.Fatal(err)
	}

	expectedRes := `{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"bWF4","signer":"cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8"}}],"sequence":"0"}`
	assert.Equal(t, expectedRes, string(res), "result is not expected json")
}
