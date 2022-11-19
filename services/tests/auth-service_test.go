package tests

import (
	"github.com/pkg/errors"
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/services"
	"github.com/rafacteixeira/calendario-medico-api/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUpWithInvalidLoginShouldReturnError(t *testing.T) {

	err := services.SignUp(util.AuthRequest{
		Login:    "",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignUpWithInvalidPwdShouldReturnError(t *testing.T) {

	err := services.SignUp(util.AuthRequest{
		Login:    "Rafa",
		Password: "",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignUpWithExistingUserShouldReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 1,
		}
	}

	err := services.SignUp(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignUpWithNewUserShouldntReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 0,
		}
	}

	services.CreateUser = func(user *model.User) {

	}

	err := services.SignUp(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.NoError(t, err, "Shouldn't return an error")
}

func TestSignInWithInvalidLoginShouldReturnError(t *testing.T) {

	_, err := services.SignIn(util.AuthRequest{
		Login:    "",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignInWithInvalidPwdShouldReturnError(t *testing.T) {

	_, err := services.SignIn(util.AuthRequest{
		Login:    "Rafa",
		Password: "",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignInWithNewUserShouldReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 0,
		}
	}

	_, err := services.SignIn(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignInErrorPasswordValidationShouldReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 1,
		}
	}

	services.ValidatePassword = func(encrypted, raw string) (bool, error) {
		return false, errors.New("error")
	}

	_, err := services.SignIn(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignInInvalidPasswordValidationShouldReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 1,
		}
	}

	services.ValidatePassword = func(encrypted, raw string) (bool, error) {
		return false, nil
	}

	_, err := services.SignIn(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.Error(t, err, "Should return an error")
}

func TestSignInValidSignInShouldntReturnError(t *testing.T) {

	services.FindUserByLogin = func(login string) model.User {
		return model.User{
			ID: 1,
		}
	}

	services.ValidatePassword = func(encrypted, raw string) (bool, error) {
		return true, nil
	}

	services.GenerateToken = func(login string) (string, error) {
		return "token", nil
	}

	_, err := services.SignIn(util.AuthRequest{
		Login:    "Rafa",
		Password: "123",
	})

	assert.NoError(t, err, "Shouldn't return an error")
}
