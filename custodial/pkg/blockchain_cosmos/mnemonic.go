// Package blockchain_cosmos defines methods to manage keys and transactions related to cosmos based chain
package blockchain_cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonic generates 24 word mnemonic
func GenerateMnemonic() (*string, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return &mnemonic, nil
}

// GetPrivKey returns private key for given mnemonic and cosmos path "m/44'/118'/0'/0/0"
func GetPrivKey(mnemonic string) (*secp256k1.PrivKey, error) {
	derivationPath := "m/44'/118'/0'/0/0"
	privKeyBytes, err := hd.Secp256k1.Derive()(mnemonic, "", derivationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from mnemonic: %w", err)
	}
	privKey := secp256k1.PrivKey{
		Key: privKeyBytes,
	}
	return &privKey, nil
}
