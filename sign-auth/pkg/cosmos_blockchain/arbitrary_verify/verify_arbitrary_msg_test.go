package arbitraryverify

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/assert"
)

func Test_VerifyArbitraryMsg(t *testing.T) {
	bytePubKey, err := hex.DecodeString("0386373dcc17c7e323cc4b2044895ce2325422abaad4270e406982cfa744667dce")
	if err != nil {
		log.Fatal("failed to get byte pubkey from hex: ", err)
	}

	hexSign := "87b00857729e8dbf4f78246fbb5dcdb706679ad979cf59b9d4a1039e2eabe41d3570432af5cfd393e01f01245115ba4d4cb8c1c43bf51a5b7b78dced8a6aa18f"
	byteSignature, err := hex.DecodeString(hexSign)
	if err != nil {
		log.Fatal("failed to get byte sign from hex: ", err)
	}
	pubKey := secp256k1.PubKey{Key: bytePubKey}
	stat, err := VerifyArbitraryMsg("cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", "max", byteSignature, pubKey)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, stat, "verify result should be true")
}
