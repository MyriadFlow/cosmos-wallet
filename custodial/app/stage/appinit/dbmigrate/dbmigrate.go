package dbmigrate

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/logo"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.CustodialUser{})
	if err != nil {
		logo.Fatalf("failed to migrate user into database: %s", err)
	}
}
