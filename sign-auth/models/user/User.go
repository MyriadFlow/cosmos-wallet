// Package user provides model and methods for storing and retriving user identified by wallet address
package user

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
)

// CustodialUser custodial user model with wallet address and one to many relation with FlowId
type User struct {
	WalletAddress string          `json:"-" gorm:"primaryKey;not null"`
	FlowId        []flowid.FlowId `gorm:"foreignkey:WalletAddress" json:"-"`
}

// Add adds user with given wallet address to database
func Add(walletAddress string) error {
	db := store.DB
	newUser := User{
		WalletAddress: walletAddress,
	}
	err := db.Model(&newUser).Create(&newUser).Error
	return err
}

// Get returns user with given wallet address from database
func Get(walletAddr string) (*User, error) {
	db := store.DB
	var user User
	res := db.Find(&user, User{
		WalletAddress: walletAddr,
	})

	if err := res.Error; err != nil {
		err = fmt.Errorf("failed to get user from database: %w", err)
		return nil, err
	}

	if res.RowsAffected == 0 {
		return nil, errorso.ErrRecordNotFound
	}

	return &user, nil
}
