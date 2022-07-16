package paseto

import (
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	pasetoclaims "github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/paseto/paseto_claims"
	"github.com/vk-rv/pvx"
)

//Returns paseto token for given wallet address
func GetPasetoForUser(walletAddr string) (string, error) {
	customClaims := pasetoclaims.New(walletAddr)
	privateKey := env.MustGetEnv("PASETO_PRIVATE_KEY")
	symK := pvx.NewSymmetricKey([]byte(privateKey), pvx.Version4)
	pv4 := pvx.NewPV4Local()
	tokenString, err := pv4.Encrypt(symK, customClaims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
