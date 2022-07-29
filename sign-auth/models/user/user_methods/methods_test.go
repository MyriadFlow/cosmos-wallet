package usermethods

import (
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Get(t *testing.T) {
	appinit.Init()

	walletAddress := "cosmos1dcspeqkh6dan5efwvs6zrjez0x0nx0s0sm6cmn"
	t.Run("Should create flow Id for new user", func(t *testing.T) {
		flowId, err := CreateFlowId(walletAddress)
		if err != nil {
			t.Fatal(err)
		}
		assert.Len(t, flowId, 36, "flowid should be of 36 charactors")
	})

	t.Run("Should create flow Id for existing user", func(t *testing.T) {
		flowId, err := CreateFlowId(walletAddress)
		if err != nil {
			t.Fatal(err)
		}
		assert.Len(t, flowId, 36, "flowid should be of 36 charactors")
	})

}
