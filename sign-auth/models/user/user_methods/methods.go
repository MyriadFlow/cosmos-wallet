package usermethods

import (
	"errors"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/errorso"
	"github.com/google/uuid"
)

//Create and insert flow Id into the database and return it
func CreateFlowId(walletAddress string) (string, error) {

	//Check if user exist
	_, err := user.Get(walletAddress)
	if err != nil {
		if errors.Is(err, errorso.ErrRecordNotFound) {
			//If doesn't exist then add that
			err = user.Add(walletAddress)
			if err != nil {
				return "", fmt.Errorf("failed to add user: %w", err)
			}
		} else {
			return "", fmt.Errorf("failed to check if user exist: %w", err)
		}
	}

	flowIdString := uuid.NewString()
	err = flowid.AddFlowId(walletAddress, flowIdString)
	if err != nil {
		return "", fmt.Errorf("failed to add flowId into database: %w", err)
	}

	return flowIdString, nil
}
