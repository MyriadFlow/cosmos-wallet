// Package paseto provides methods to generate and verify paseto tokens
package paseto

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	pasetoclaims "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto/paseto_claims"
	"github.com/vk-rv/pvx"
)

//Returns paseto token for given wallet address
func GetPasetoForUser(walletAddr string) (string, error) {
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
	expiration := pasetoExpirationInHoursInt * time.Hour
	signedBy := env.MustGetEnv("SIGNED_BY")
	customClaims := pasetoclaims.New(walletAddr, expiration, signedBy)
	privateKey := env.MustGetEnv("PASETO_PRIVATE_KEY")
	symK := pvx.NewSymmetricKey([]byte(privateKey), pvx.Version4)
	pv4 := pvx.NewPV4Local()
	tokenString, err := pv4.Encrypt(symK, customClaims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyPaseto(pasetoToken string) error {
	pv4 := pvx.NewPV4Local()
	k := env.MustGetEnv("PASETO_PRIVATE_KEY")
	symK := pvx.NewSymmetricKey([]byte(k), pvx.Version4)
	var cc pasetoclaims.CustomClaims
	err := pv4.
		Decrypt(pasetoToken, symK).
		ScanClaims(&cc)
	if err != nil {
		err = fmt.Errorf("failed to scan claims: %w", err)
		return err
	}
	return nil
}
