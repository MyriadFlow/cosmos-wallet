package appinit

import (
	"github.com/MyriadFlow/cosmos-wallet/app/stage/appinit/dbconinit"
	"github.com/MyriadFlow/cosmos-wallet/app/stage/appinit/dbmigrate"
	"github.com/MyriadFlow/cosmos-wallet/app/stage/appinit/logoinit"
)

func Init() {
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
