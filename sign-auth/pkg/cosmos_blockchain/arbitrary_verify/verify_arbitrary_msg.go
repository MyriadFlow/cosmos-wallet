package arbitraryverify

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// VerifyArbitraryMsg verifies arbitrary Adr036 message by first
// composing it with the 0 values required and
// then verifying it against public key
func VerifyArbitraryMsg(signer string, msg string, signature []byte, publicKey secp256k1.PubKey) (bool, error) {
	composedArbitraryMsg, err := ComposeArbitraryMsg(signer, msg)
	if err != nil {
		return false, fmt.Errorf("failed to compose arbitrary msg: %w", err)
	}

	verifyResult := publicKey.VerifySignature(composedArbitraryMsg, signature)
	return verifyResult, nil
}
