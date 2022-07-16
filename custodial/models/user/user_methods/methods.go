package usermethods

import (
	"errors"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"

	"github.com/google/uuid"
)

// Creates and stores mnemonic and returns user Id
func Create() (string, error) {
	mnemonic, err := blockchain_cosmos.GenerateMnemonic()
	if err != nil {
		return "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}
	uid := uuid.NewString()

	err = user.Add(uid, *mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to add user into database: %w", err)
	}

	return uid, nil
}

func Transfer(uid string, from string, to string, amount int64) error {
	return errors.New("not implemented")
}
