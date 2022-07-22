package main

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/apprun"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appinit.Init()
	logo.Info("Starting Custodial Wallet")
	apprun.Run()
}
