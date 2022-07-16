package main

import (
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
)

func main() {
	appinit.Init()
	logo.Info("Hola Cosmos")
}
