package environment

import (
	"log"

	"github.com/MyriadFlow/cosmos-wallet/pkg/env"
)

type Environment int

const (
	PROD Environment = iota
	DEV  Environment = iota
)

func GetEnvironment() Environment {
	appEnv := env.MustGetEnv("APP_ENVIRONMENT")

	if appEnv == "PROD" {
		return PROD
	} else if appEnv == "DEV" {
		return DEV
	} else {
		log.Fatal("App environment not supported")
		return -1
	}
}
