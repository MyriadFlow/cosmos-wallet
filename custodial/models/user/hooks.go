package user

import (
	"encoding/hex"
	"strings"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"gorm.io/gorm"
)

// BeforeSave encodes the mnemonic in hexadecimal before saving it into database
func (u *CustodialUser) BeforeSave(tx *gorm.DB) (err error) {
	hexMnemonic := "0x" + hex.EncodeToString([]byte(u.Mnemonic))
	u.Mnemonic = hexMnemonic
	return nil
}

// AfterFind decodes the mnemonic from hexadecimal after finding it from database
func (u *CustodialUser) AfterFind(tx *gorm.DB) (err error) {
	hexStringWithout0x := strings.TrimPrefix(u.Mnemonic, "0x")
	plainMnemonic, err := hex.DecodeString(hexStringWithout0x)
	if err != nil {
		logo.Errorf("AfterFind: failed to decode mnemonic from hex")
		return err
	}
	u.Mnemonic = string(plainMnemonic)
	return nil
}
