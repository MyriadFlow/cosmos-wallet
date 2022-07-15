package user

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/models/flowid"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/store"
)

type User struct {
	WalletAddress string          `json:"-" gorm:"primaryKey;not null"`
	FlowId        []flowid.FlowId `gorm:"foreignkey:WalletAddress" json:"-"`
}

func Add(walletAddress string) error {
	db := store.DB
	newUser := User{
		WalletAddress: walletAddress,
	}
	if err := db.Model(&newUser).Create(&newUser).Error; err != nil {
		return err
	} else {
		return nil
	}
}

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
