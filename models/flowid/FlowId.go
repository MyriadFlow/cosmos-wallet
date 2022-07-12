package flowid

import "github.com/MyriadFlow/cosmos-wallet/pkg/store"

type FlowId struct {
	WalletAddress string
	FlowId        string `gorm:"primary_key"`
}

func GetFlowId(flowId string) (*FlowId, error) {
	db := store.DB
	var userFlowId FlowId
	err := db.Find(&userFlowId, &FlowId{
		FlowId: flowId,
	}).Error

	if err != nil {
		return nil, err
	}
	return &userFlowId, err
}
