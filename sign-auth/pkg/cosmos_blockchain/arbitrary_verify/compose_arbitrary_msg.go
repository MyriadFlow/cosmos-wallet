package arbitraryverify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
)

// ComposeArbitraryMsg Creates SignDoc with JSON encoded bytes as per adr036
// Compatible with AMINO as it is supported by keplr wallet
func ComposeArbitraryMsg(signer string, data string) ([]byte, error) {
	data_base64 := base64.StdEncoding.EncodeToString([]byte(data))

	newSignDocMsgValue := signDocMsgValue{
		Data:   data_base64,
		Signer: signer,
	}

	newSignDocMsg := signDocMsg{
		Value: newSignDocMsgValue, Type: "sign/MsgSignData",
	}
	newSignDoc := signDoc{
		Msgs: []signDocMsg{
			newSignDocMsg,
		},
		AccountNumber: "0",
		Sequence:      "0",
		Fee: signDocFee{
			Gas:    "0",
			Amount: types.NewCoins(),
		},
	}

	jsonBytes, err := json.Marshal(newSignDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to Sign Doc to JSON: %w", err)
	}
	return jsonBytes, nil
}
