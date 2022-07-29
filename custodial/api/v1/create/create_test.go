package create

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostCreate(t *testing.T) {
	appinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())

	_ = "/api/v1.0/create"
	t.Run("Should return 200 and UserId", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		create(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode, "status code not 200 (OK), body: %s", rr.Body)
	})

}
