package flowidmethods

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid"
	arbitraryverify "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/cosmos_blockchain/arbitrary_verify"
	walletaddress "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/cosmos_blockchain/wallet_address"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

var ErrSignDenied = errors.New("signature denied")
var ErrPubKeyDenied = errors.New("public key denied")

// VerifySignAndGetPaseto verifies the signature for given flowID and returns paseto if it is valid
//
// Also deletes the flow id after approving signature
func VerifySignAndGetPaseto(publicKey secp256k1.PubKey, signatureBase64 string, flowId string) (string, error) {

	dataFlowId, err := flowid.GetFlowId(flowId)
	if err != nil {
		return "", fmt.Errorf("failed to get flow id from database: %w", err)
	}

	// Prepare expected signing data (msg)
	authEula := env.MustGetEnv("AUTH_EULA")
	signingData := fmt.Sprintf("%s%s", authEula, dataFlowId.FlowId)
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 signature: %w", err)
	}

	//Check if public key is of same wallet address as of flow ID
	signerFromPubKey, err := walletaddress.GetWalletAddrFromPubKey(env.MustGetEnv("WALLET_ADDRESS_HRP"), publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to get signer from public key: %w", err)
	}

	if signerFromPubKey != dataFlowId.WalletAddress {
		return "", ErrPubKeyDenied
	}
	signatureApproved, err := arbitraryverify.VerifyArbitraryMsg(dataFlowId.WalletAddress, signingData, signatureBytes, publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to verify arbitrary signature: %w", err)
	}

	//If signature not approved then return error
	if !signatureApproved {
		return "", ErrSignDenied
	}

	paseto, err := paseto.GetPasetoForUser(dataFlowId.WalletAddress)
	if err != nil {
		return "", fmt.Errorf("failed to generate paseto: %w", err)
	}

	err = flowid.DeleteFlowId(flowId)
	if err != nil {
		return "", fmt.Errorf("failed to delete flowid: %w", err)
	}
	return paseto, nil
}
