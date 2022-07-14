package blockchain_cosmos

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
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

func GetWallet(mnemonic string) ([]byte, error) {
	return hd.Secp256k1.Derive()(mnemonic, "", types.FullFundraiserPath)
}
