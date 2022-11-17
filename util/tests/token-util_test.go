package utiltests

import (
	"github.com/rafacteixeira/calendario-medico-api/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIfAValidTokenIsConsideredValid(t *testing.T) {

	token, _ := util.GenerateToken("rafael")
	valid, _ := util.ValidateAdminToken(token)

	assert.Equal(t, true, valid, "Should be valid")
}

func TestIfAnInvalidTokenIsConsideredInvalid(t *testing.T) {

	token, _ := util.GenerateToken("rafael")
	valid, _ := util.ValidateAdminToken(token + "a")

	assert.Equal(t, false, valid, "Should be invalid")
}

func TestIfAnExpiredTokenIsConsideredInvalid(t *testing.T) {

	util.GetTokenExpirationFunc = func() time.Time {
		return time.Now().Add(time.Second * 1)
	}

	token, _ := util.GenerateToken("rafael")
	time.Sleep(time.Second * 2)
	valid, _ := util.ValidateAdminToken(token)

	assert.Equal(t, false, valid, "Should be invalid")
}
