package services

import (
	"github.com/pkg/errors"
	"github.com/rafacteixeira/calendario-medico-api/database"
	"github.com/rafacteixeira/calendario-medico-api/util"
)

var (
	FindUserByLogin  = database.FindUser
	CreateUser       = database.CreateUser
	ValidatePassword = util.Validate
	GenerateToken    = util.GenerateToken
)

func SignUp(request util.AuthRequest) error {

	if len(request.Login) > 0 && len(request.Password) > 0 {

		user := FindUserByLogin(request.Login)
		if user.ID != 0 {
			return errors.New("Login already registered")
		}

		encryptedPassword, err := util.Encrypt(request.Password)
		if err != nil {
			return err
		}

		user.Login = request.Login
		user.Password = encryptedPassword
		CreateUser(&user)
	} else {
		return errors.New("Invalid User/Password combination!")
	}

	return nil
}

func SignIn(request util.AuthRequest) (string, error) {

	if len(request.Login) > 0 && len(request.Password) > 0 {

		user := FindUserByLogin(request.Login)
		if user.ID == 0 {
			return "", errors.New("User Not Found")
		}

		validPwd, err := ValidatePassword(user.Password, request.Password)
		if err != nil {
			return "", err
		}
		if !validPwd {
			return "", errors.New("Invalid Password!")
		}

		return GenerateToken(request.Login)

	} else {
		return "", errors.New("Invalid User/Password combination!")
	}

}

func CheckToken(token string) (util.UserClaims, error) {
	_, err, claims := util.ValidateToken(token)
	return claims, err

}
