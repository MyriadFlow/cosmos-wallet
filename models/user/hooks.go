package user

import (
	"encoding/hex"

	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hexMnemonic := "0x" + hex.EncodeToString([]byte(u.Mnemonic))
	u.Mnemonic = hexMnemonic
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	plainMnemonic, err := hex.DecodeString(string(u.Mnemonic[2:]))
	if err != nil {
		return err
	}
	u.Mnemonic = string(plainMnemonic)
	return nil
}
