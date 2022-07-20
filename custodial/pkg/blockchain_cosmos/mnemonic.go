package blockchain_cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tyler-smith/go-bip39"
)

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

//TODO: confirm it - Returns private key for given mnemonic and path in cosmos sdk ("github.com/cosmos/cosmos-sdk/types").FullFundraiserPath
func GetWallet(mnemonic string) (*secp256k1.PrivKey, error) {
	privKeyBytes, err := hd.Secp256k1.Derive()(mnemonic, "", types.FullFundraiserPath)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from mnemonic: %w", err)
	}
	privKey := secp256k1.PrivKey{
		Key: privKeyBytes,
	}
	return &privKey, nil
}
