package transfer

type TransferRequest struct {
	UserId string `json:"userId"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

type TransferPayload struct {
	TransactionHash string `json:"hash"`
}
