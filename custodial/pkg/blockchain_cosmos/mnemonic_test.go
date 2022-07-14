package blockchain_cosmos

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetWalletAddr(t *testing.T) {
	mnemonic := "regret notable delay giraffe trip surge icon comfort maple swift bounce spy maze side apology van top mercy dice lesson remain regular coast pony"
	expectedHexPrivKey := "e85362ff644e9ea00d633c8cd92e51c0b9cfbd7f679975ccadaf86d9dc43fd29"
	privKey, err := GetWallet(mnemonic)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedHexPrivKey, hex.EncodeToString(privKey))
}
