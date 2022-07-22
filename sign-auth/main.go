package main

import (
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/apprun"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appinit.Init()
	logo.Info("Starting Sign Auth")
	apprun.Run()
}
