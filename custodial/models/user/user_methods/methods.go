package usermethods

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/google/uuid"
)

// Creates and stores mnemonic and returns public Key and user Id
func Create() (*cryptotypes.PubKey, string, error) {
	mnemonic, err := blockchain_cosmos.GenerateMnemonic()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}
	uid := uuid.NewString()

	err = user.Add(uid, *mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to add user into database: %w", err)
	}

	privKey, err := blockchain_cosmos.GetWallet(*mnemonic)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get privateKey for mnemonic: %w", err)
	}
	pubKey := privKey.PubKey()
	return &pubKey, uid, nil
}

// Creates transfer request for user using private key, returns transaction hash and error if any
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
	privKey, err := blockchain_cosmos.GetWallet(user.Mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to get private key: %w", err)
	}

	hash, err := blockchain_cosmos.Transfer(&blockchain_cosmos.TransferParams{
		FromAddr: fromAddr,
		ToAddr:   toAddr,
		PrivKey:  privKey,
		Denom:    "uatom",
		Amount:   amount,
	})

	if err != nil {
		err = fmt.Errorf("failed to transfer amount: %w", err)
		return "", err
	}

	return hash, err
}
