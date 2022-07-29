// Package walletaddress provides method to get wallet address from public key
package walletaddress

import (
	"crypto/sha256"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"golang.org/x/crypto/ripemd160"
)

// GetWalletAddrFromPubKey returns bech32 wallet address from secp256k1.PubKey and with provided hrp
func GetWalletAddrFromPubKey(hrp string, publicKey secp256k1.PubKey) (string, error) {
	pubKey_shasha256 := sha256.Sum256(publicKey.Bytes())
	pubKey_shasha256_ripemd160 := ripemd160.New()
	pubKey_shasha256_ripemd160.Write(pubKey_shasha256[:])
	rip := pubKey_shasha256_ripemd160.Sum(nil)
	walletAddr, err := bech32.ConvertAndEncode(hrp, rip)
	if err != nil {
		return "", fmt.Errorf("failed to convert and encode ripemd160 to bech32: %w", err)
	}
	return walletAddr, nil
}
