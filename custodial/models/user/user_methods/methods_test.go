package usermethods

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/custodial/models/user"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/blockchain_cosmos"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/store"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/testingcommon"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Get(t *testing.T) {
	appinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	var (
		uid          string
		err          error
		base64PubKey string
	)
	t.Run("create user", func(t *testing.T) {
		pubKey, _uid, err := Create()
		if err != nil {
			t.Fatal(err)
		}
		uid = _uid

		base64PubKey = base64.StdEncoding.EncodeToString((*pubKey).Bytes())
		assert.Len(t, _uid, 36)
		assert.Len(t, base64PubKey, 44)
	})

	t.Run("get user", func(t *testing.T) {
		fetchedUser, err := user.Get(uid)
		if err != nil {
			t.Fatal(err)
		}
		privKey, err := blockchain_cosmos.GetPrivKey(fetchedUser.Mnemonic)
		_base64PubKey := base64.StdEncoding.EncodeToString(privKey.PubKey().Bytes())
		assert.Equal(t, base64PubKey, _base64PubKey)
		assert.Equal(t, fetchedUser.Id, uid)
	})

	t.Run("transfer atom", func(t *testing.T) {
		uid = uuid.NewString()
		// mnemonic with balance for testing transfer
		mnemonic := "envelope rebel nerve sock change animal such hero pave bomb coffee invest misery detect enhance muffin stable bundle ski equal have shadow seed arena"
		//Clean before since the wallet address is not generated
		db := store.DB
		hexMnemonic := "0x" + hex.EncodeToString([]byte(mnemonic))
		delRes := db.Where("mnemonic = ?", hexMnemonic).Delete(&user.CustodialUser{})
		if delRes.Error != nil {
			t.Fatal(err)
		}

		user.Add(uid, mnemonic)
		_, err := Transfer(uid, "cosmos1fzqqen9f9jwsc6x5v7hltdm4ctxhvpdvna8n3p", "cosmos1uuyak34fv767a65k9f4ms8jepcc2z5wswt5eg8", 1)
		if err != nil {
			t.Fatal(err)
		}
	})
}
