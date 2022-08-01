package create

type CreatePayload struct {
	UserId     string `json:"userId"`
	PublicKey  string `json:"publicKey"`
	WalletAddr string `json:"walletAddr"`
}
