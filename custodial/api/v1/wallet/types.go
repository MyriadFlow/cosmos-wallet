package wallet

type GetWalletRequest struct {
	UserId string `json:"userId"`
}

type GetWalletPayload struct {
	PublicKey  string `json:"publicKey"`
	WalletAddr string `json:"walletAddr"`
}
