package main

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appinit.Init()
	logo.Info("Hello Cosmos")
}
