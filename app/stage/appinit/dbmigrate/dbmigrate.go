package dbmigrate

import (
	"github.com/MyriadFlow/cosmos-wallet/models/user"
	"github.com/MyriadFlow/cosmos-wallet/pkg/logo"
	"github.com/MyriadFlow/cosmos-wallet/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		logo.Fatalf("failed to migrate user into database: %s", err)
	}
}
