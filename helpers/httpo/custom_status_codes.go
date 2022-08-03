package httpo

const (
	// Auth issues

	// Occurs when token is expired
	TokenExpired = 4031

	// Occurs when token is invalid, for example, signed by wrong signature or is malformed
	TokenInvalid = 4033

	// Occurs when signatures public key doesn't match to the one which was used while requesting challenge
	SignatureDenied = 4034

	// Requet params issues

	// The header doesn't contain Authorization header or it is empty
	AuthHeaderMissing = 4001

	// The provided string is not valid base64
	InvalidBase64 = 4002

	// The provided wallet address is not compatible to the chain
	WalletAddressInvalid = 4003

	// State issues

	// User trying to refer doesn't exist in database
	UserNotFound = 4041

	// FlowID trying to refer doesn't exist in database
	FlowIdNotFound = 4042

	// Account trying to refer in chain doesn't exist, this means that account doesn't have any in or out transactions and therefore it has 0 balance
	AccountNotFound = 4043
)
