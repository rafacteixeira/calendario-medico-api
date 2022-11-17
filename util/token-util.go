package util

import (
	"github.com/cristalhq/jwt/v4"
	"github.com/rafacteixeira/calendario-medico-api/settings"
	"time"
)

var key = []byte(settings.TokenSecretSeed())
var GetTokenExpirationFunc = getTokenExpiration

func GenerateToken(login string) (string, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	if err != nil {
		return "", err
	}

	expiresAt := GetTokenExpirationFunc()

	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: []string{"admin"},
			ID:       "random-unique-string",
			ExpiresAt: &jwt.NumericDate{
				Time: expiresAt,
			},
		},
		Login: login,
	}

	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func getTokenExpiration() time.Time {
	return time.Now().Add(time.Hour * time.Duration(1))
}

func ValidateAdminToken(tokenStr string) (bool, error) {

	key := []byte(settings.TokenSecretSeed())
	verifier, verifierError := jwt.NewVerifierHS(jwt.HS256, key)
	if verifierError != nil {
		return false, verifierError
	}

	tokenBytes := []byte(tokenStr)

	var newClaims jwt.RegisteredClaims
	parseClaimsError := jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	if parseClaimsError != nil {
		return false, parseClaimsError
	}

	verifyAdmin := newClaims.IsForAudience("admin")
	verifyExpiration := newClaims.IsValidAt(time.Now())

	return verifyAdmin && verifyExpiration, nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	Login string
}
