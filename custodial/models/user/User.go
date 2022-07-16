package user

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/store"
)

type CustodialUser struct {
	Id       string `json:"-" gorm:"primaryKey;not null"`
	Mnemonic string `json:"-" gorm:"unique;not null"`
}

func Add(id string, mnemonic string) error {
	db := store.DB
	newUser := CustodialUser{
		Id:       id,
		Mnemonic: mnemonic,
	}
	err := db.Model(&newUser).Create(&newUser).Error
	return err
}

func Get(id string) (*CustodialUser, error) {
	db := store.DB
	var user CustodialUser
	res := db.Find(&user, CustodialUser{
		Id: id,
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
