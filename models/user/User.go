package user

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/pkg/errorso"
	"github.com/MyriadFlow/cosmos-wallet/pkg/store"
)

type User struct {
	Id       string `json:"-" gorm:"primaryKey;not null"`
	Mnemonic string `json:"-" gorm:"unique;not null"`
}

func Add(id string, mnemonic string) error {
	db := store.DB
	newUser := User{
		Id:       id,
		Mnemonic: mnemonic,
	}
	if err := db.Model(&newUser).Create(&newUser).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func Get(id string) (*User, error) {
	db := store.DB
	var user User
	res := db.Find(&user, User{
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
