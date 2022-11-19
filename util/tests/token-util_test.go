package utiltests

import (
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIfAValidTokenIsConsideredValid(t *testing.T) {

	util.GetTokenSecret = func() string {
		return "test"
	}

	util.GetUserWithRoles = func(login string) model.User {
		return model.User{}
	}

	token, _ := util.GenerateToken("rafael")
	valid, _, _ := util.ValidateToken(token)

	assert.Equal(t, true, valid, "Should be valid")
}

func TestIfAnInvalidTokenIsConsideredInvalid(t *testing.T) {

	util.GetTokenSecret = func() string {
		return "test"
	}

	token, _ := util.GenerateToken("rafael")
	valid, _, _ := util.ValidateToken(token + "a")

	assert.Equal(t, false, valid, "Should be invalid")
}

func TestIfAnExpiredTokenIsConsideredInvalid(t *testing.T) {

	util.GetTokenSecret = func() string {
		return "test"
	}

	util.GetTokenExpirationFunc = func() time.Time {
		return time.Now().Add(time.Second * 1)
	}

	token, _ := util.GenerateToken("rafael")
	time.Sleep(time.Second * 2)
	valid, _, _ := util.ValidateToken(token)

	assert.Equal(t, false, valid, "Should be invalid")
}
