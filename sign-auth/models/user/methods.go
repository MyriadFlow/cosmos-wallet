package user

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
	"github.com/google/uuid"
)

//Create and insert flow Id into the database and return it
func CreateFlowId(walletAddress string) (string, error) {
	db := store.DB
	association := db.Model(&User{WalletAddress: walletAddress}).Association("FlowId")
	flowIdString := uuid.NewString()
	newFlowId := flowid.FlowId{
		WalletAddress: walletAddress,
		FlowId:        flowIdString,
	}
	err := association.Append(&newFlowId)
	if err != nil {
		return "", fmt.Errorf("failed to add flowId into database: %w", err)
	}

	return flowIdString, nil
}
