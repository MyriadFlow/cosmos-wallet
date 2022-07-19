package testingcommon

import (
	"time"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	pasetoclaims "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto/paseto_claims"
	"github.com/vk-rv/pvx"
)

//Returns paseto token for given wallet address and with expiration, only use in tests
func GetPasetoForTestUser(walletAddr string, expiration time.Duration) (string, error) {
	customClaims := pasetoclaims.New(walletAddr, expiration, env.MustGetEnv("SIGNED_BY"))
	privateKey := env.MustGetEnv("PASETO_PRIVATE_KEY")
	symK := pvx.NewSymmetricKey([]byte(privateKey), pvx.Version4)
	pv4 := pvx.NewPV4Local()
	tokenString, err := pv4.Encrypt(symK, customClaims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
