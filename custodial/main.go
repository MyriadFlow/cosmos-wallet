package main

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
)

func main() {
	appinit.Init()
	logo.Info("Hello Cosmos")
}
