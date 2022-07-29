// Package env provides MustGetEnv to get env
package env

import (
	"os"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
)

// MustGetEnv returns value of env variable if it exist or if it doesn't exist
// then fatal logs the key
func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logo.Fatalf("env variable %v is not defined", key)
	}
	return val
}
