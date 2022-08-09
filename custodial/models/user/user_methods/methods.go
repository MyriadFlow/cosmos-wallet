// Package usermethods provides core methods for user logic
package usermethods

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos/cosmos_transfer"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/google/uuid"
)

// Create creates and stores mnemonic and returns public Key and user Id
func Create() (*cryptotypes.PubKey, string, error) {

	// Generate mnemonic
	mnemonic, err := blockchain_cosmos.GenerateMnemonic()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}

	// Generate user id
	uid := uuid.NewString()

	err = user.Add(uid, *mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to add user into database: %w", err)
	}

	// Get private key from mnemonic
	privKey, err := blockchain_cosmos.GetPrivKey(*mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get privateKey for mnemonic: %w", err)
	}

	pubKey := privKey.PubKey()
	return &pubKey, uid, nil
}

// Transfer creates transfer request for user using private key, returns transaction hash and error if any
func Transfer(uid string, from string, to string, amount int64) (txHash string, erR error) {
	user, err := user.Get(uid)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return "", fmt.Errorf("failed to get from address: %w", err)
	}

	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return "", fmt.Errorf("failed to get to address: %w", err)
	}

	// Get private key from mnemonic
	privKey, err := blockchain_cosmos.GetPrivKey(user.Mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to get private key: %w", err)
	}
	// Create transfer request
	hash, err := cosmos_transfer.Transfer(&cosmos_transfer.TransacParams{
		FromAddr: fromAddr,
		ToAddr:   toAddr,
		PrivKey:  privKey,
		Denom:    env.MustGetEnv("SMALLEST_DENOM"),
		Amount:   amount,
	})

	if err != nil {
		err = fmt.Errorf("failed to transfer amount: %w", err)
		return "", err
	}

	// Return transaction hash
	return hash, nil
}
