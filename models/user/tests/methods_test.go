package tests

import (
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/models/user"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Get(t *testing.T) {
	appinit.Init()
	var (
		uid string
		err error
	)
	t.Run("create user", func(t *testing.T) {
		uid, err = user.Create()
		if err != nil {
			t.Fatal(err)
		}
		assert.Len(t, uid, 36)
	})

	t.Run("get user", func(t *testing.T) {
		fetchedUser, err := user.Get(uid)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fetchedUser.Id, uid)
	})

	t.Run("transfer atom", func(t *testing.T) {
		err = user.Transfer(uid, "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", 1)
		if err != nil {
			t.Fatal(err)
		}
	})
}
