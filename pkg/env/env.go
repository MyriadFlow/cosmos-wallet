package env

import (
	"os"

	"github.com/MyriadFlow/cosmos-wallet/pkg/logo"
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logo.Fatalf("env variable %v is not defined", key)
	}
	return val
}
