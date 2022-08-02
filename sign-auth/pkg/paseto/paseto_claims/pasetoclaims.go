// Package pasetoclaims provides claim declaration for token generation and verification
package pasetoclaims

import (
	"time"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
	"github.com/vk-rv/pvx"
)

// CustomClaims defines claims for paseto containing wallet address, signed by and RegisteredClaims
type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	SignedBy      string `json:"signedBy"`
	pvx.RegisteredClaims
}

// Valid checks if the claims are valid agaist RegisteredClaims and checks if wallet address
// exist in database
func (c CustomClaims) Valid() error {
	db := store.DB
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	err := db.Model(&user.User{}).Where("wallet_address = ?", c.WalletAddress).First(&user.User{}).Error
	return err
}

// New returns CustomClaims with wallet address, signed by and expiration
func New(walletAddress string, expiration time.Duration, signedBy string) CustomClaims {
	expirationTime := time.Now().Add(expiration)
	return CustomClaims{
		walletAddress,
		signedBy,
		pvx.RegisteredClaims{
			Expiration: &expirationTime,
		},
	}
}
