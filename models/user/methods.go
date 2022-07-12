package user

import (
	"errors"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/pkg/blockchain_cosmos"

	"github.com/google/uuid"
)

func Create() (uid string, err error) {
	mnemonic, err := blockchain_cosmos.GenerateMnemonic()
	if err != nil {
		return "", fmt.Errorf("failed to generate mnemonic: %w", err)
	}
	uid = uuid.NewString()

	err = Add(uid, *mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to add user into database: %w", err)
	}

	return uid, nil
}

func Transfer(uid string, from string, to string, amount int64) error {
	return errors.New("not implemented")
}
