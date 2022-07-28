package httpo

//TODO custom status code and its documentation
const (
	// Auth issues
	TokenExpired    = 4031
	TokenInvalid    = 4033
	SignatureDenied = 4034

	// Requet params issues
	AuthHeaderMissing    = 4001
	InvalidBase64        = 4002
	WalletAddressInvalid = 4003

	// State issues
	UserNotFound = 4041
)
