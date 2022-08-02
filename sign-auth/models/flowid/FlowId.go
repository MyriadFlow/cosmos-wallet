// Package flowid provides model and methods for storing and retriving flowid for wallet address
package flowid

import (
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
)

type FlowId struct {
	WalletAddress string
	FlowId        string `gorm:"primary_key"`
}

func GetFlowId(flowId string) (*FlowId, error) {
	db := store.DB
	var userFlowId FlowId
	res := db.Find(&userFlowId, &FlowId{
		FlowId: flowId,
	})

	if err := res.Error; err != nil {
		return nil, err
	}

	if res.RowsAffected == 0 {
		return nil, errorso.ErrRecordNotFound
	}
	return &userFlowId, nil
}

//Adds flow id into database for given wallet Address
func AddFlowId(walletAddr string, flowId string) error {
	db := store.DB
	err := db.Create(&FlowId{
		WalletAddress: walletAddr,
		FlowId:        flowId,
	}).Error

	return err
}

func DeleteFlowId(flowId string) error {
	db := store.DB
	err := db.Delete(&FlowId{
		FlowId: flowId,
	}).Error

	return err
}
