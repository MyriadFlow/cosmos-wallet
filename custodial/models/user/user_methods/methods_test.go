package usermethods

import (
	"encoding/base64"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Get(t *testing.T) {
	appinit.Init()
	var (
		uid    string
		err    error
		pubKey *cryptotypes.PubKey
	)
	t.Run("create user", func(t *testing.T) {
		pubKey, uid, err = Create()
		if err != nil {
			t.Fatal(err)
		}

		base64PubKey := base64.StdEncoding.EncodeToString((*pubKey).Bytes())
		assert.Len(t, uid, 36)
		assert.Len(t, base64PubKey, 44)
	})

	t.Run("get user", func(t *testing.T) {
		fetchedUser, err := user.Get(uid)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fetchedUser.Id, uid)
	})

	t.Run("transfer atom", func(t *testing.T) {
		err = Transfer(uid, "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", 1)
		if err != nil {
			t.Fatal(err)
		}
	})
}
