package utiltests

import (
	"github.com/rafacteixeira/calendario-medico-api/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfPasswordWithNoCharsShouldntBeGenerated(t *testing.T) {

	_, err := util.Encrypt("")
	assert.Error(t, err)
	assert.Equal(t, "Password length cannot be 0", err.Error(), "Error generating password")
}

func TestIfPasswordWithCharsShouldBeGenerated(t *testing.T) {

	enc, err := util.Encrypt("123")
	assert.NoError(t, err)
	assert.Equal(t, true, len(enc) > 0, "Password should have more chars")
}

func TestValidPasswordVerification(t *testing.T) {

	enc, _ := util.Encrypt("123")
	validate, err := util.Validate(enc, "123")
	assert.NoError(t, err)
	assert.Equal(t, true, validate, "Password should be valid")
}

func TestInvalidPasswordVerification(t *testing.T) {

	enc, _ := util.Encrypt("1234")
	validate, err := util.Validate(enc, "123")
	assert.Error(t, err)
	assert.Equal(t, false, validate, "Password should be invalid")
}
