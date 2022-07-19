package pasetoclaims

import (
	"time"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
	"github.com/vk-rv/pvx"
)

type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	SignedBy      string `json:"signedBy"`
	pvx.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	db := store.DB
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	err := db.Model(&user.User{}).Where("wallet_address = ?", c.WalletAddress).First(&user.User{}).Error
	return err
}

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
