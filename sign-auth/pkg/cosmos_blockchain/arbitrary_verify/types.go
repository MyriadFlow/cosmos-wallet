// Package arbitraryverify provides methods to verify signatures signed with Arbitrary message
package arbitraryverify

import sdk "github.com/cosmos/cosmos-sdk/types"

type signDocFee struct {
	Amount []sdk.Coin `json:"amount"`
	Gas    string     `json:"gas"`
}

type signDocMsgValue struct {
	Data   string `json:"data"`
	Signer string `json:"signer"`
}

type signDocMsg struct {
	Type  string          `json:"type"`
	Value signDocMsgValue `json:"value"`
}
type signDoc struct {
	AccountNumber string       `json:"account_number"`
	ChainId       string       `json:"chain_id"`
	Fee           signDocFee   `json:"fee"`
	Memo          string       `json:"memo"`
	Msgs          []signDocMsg `json:"msgs"`
	Sequence      string       `json:"sequence"`
}
