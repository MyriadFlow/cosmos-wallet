package main

import (
	"github.com/MyriadFlow/cosmos-wallet/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/pkg/logo"
)

func main() {
	appinit.Init()
	logo.Info("Hello Cosmos")
}
