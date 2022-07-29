// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/store"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.CustodialUser{})
	if err != nil {
		logo.Fatalf("failed to migrate user into database: %s", err)
	}
}
