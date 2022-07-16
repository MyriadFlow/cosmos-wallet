package pasetoclaims

import (
	"os"
	"strconv"
	"time"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/user"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
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

func New(walletAddress string) CustomClaims {
	pasetoExpirationInHours, ok := os.LookupEnv("PASETO_EXPIRATION_IN_HOURS")
	pasetoExpirationInHoursInt := time.Duration(24)
	if ok {
		res, err := strconv.Atoi(pasetoExpirationInHours)
		if err != nil {
			logo.Warnf("Failed to parse PASETO_EXPIRATION_IN_HOURS as int : %v", err.Error())
		} else {
			pasetoExpirationInHoursInt = time.Duration(res)
		}
	}
	expiration := time.Now().Add(pasetoExpirationInHoursInt * time.Hour)
	signedBy := env.MustGetEnv("SIGNED_BY")
	return CustomClaims{
		walletAddress,
		signedBy,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}
