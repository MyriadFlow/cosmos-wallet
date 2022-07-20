package usermethods

import (
	"errors"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

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

func Transfer(uid string, from string, to string, amount int64) error {
	return errors.New("not implemented")
}
