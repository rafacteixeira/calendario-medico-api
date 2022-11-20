package util

import (
	"github.com/cristalhq/jwt/v4"
	"github.com/google/uuid"
	"github.com/rafacteixeira/calendario-medico-api/database"
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/settings"
	"time"
)

var GetTokenSecret = settings.TokenSecretSeed
var GetTokenExpirationFunc = getTokenExpiration
var GetUserWithRoles = database.FindUserWithRoles
var SaveTokenInfo = database.CreateToken
var GetTokenInfo = database.FindToken
var DeleteToken = database.DeleteToken

func GenerateToken(login string) (string, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(GetTokenSecret()))

	if err != nil {
		return "", err
	}

	expiresAt := GetTokenExpirationFunc()
	user := GetUserWithRoles(login)

	jti := uuid.New().String()
	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: []string{"admin"},
			ID:       jti,
			ExpiresAt: &jwt.NumericDate{
				Time: expiresAt,
			},
		},
		UserLogin: user.Login,
		UserID:    user.ID,
		UserRoles: user.Roles,
	}
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	if err != nil {
		return "", err
	}

	tokenStr := token.String()
	SaveTokenInfo(model.Token{JTI: jti, Token: tokenStr})
	return tokenStr, nil
}

func getTokenExpiration() time.Time {
	return time.Now().Add(time.Hour * 24 * time.Duration(30))
}

func ValidateToken(tokenStr string) (bool, error, UserClaims) {
	verifier, verifierError := jwt.NewVerifierHS(jwt.HS256, []byte(GetTokenSecret()))
	if verifierError != nil {
		return false, verifierError, UserClaims{}
	}

	tokenBytes := []byte(tokenStr)

	var newClaims UserClaims
	parseClaimsError := jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	if parseClaimsError != nil {
		return false, parseClaimsError, newClaims
	}

	token := GetTokenInfo(newClaims.ID)
	verifyExisting := token != model.Token{}
	verifyAudience := newClaims.IsForAudience("admin")
	verifyExpiration := newClaims.IsValidAt(time.Now())

	return verifyExisting && verifyAudience && verifyExpiration, nil, newClaims
}

func ValidateRole(tokenStr string, role string) (bool, error) {

	verifier, verifierError := jwt.NewVerifierHS(jwt.HS256, []byte(GetTokenSecret()))
	if verifierError != nil {
		return false, verifierError
	}

	tokenBytes := []byte(tokenStr)

	var newClaims UserClaims
	parseClaimsError := jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	if parseClaimsError != nil {
		return false, parseClaimsError
	}

	return newClaims.HasRole(role), nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserLogin string
	UserID    uint
	UserRoles []model.Role
}

func (claims UserClaims) HasRole(roleToCheck string) bool {
	for _, role := range claims.UserRoles {
		if role.Name == roleToCheck {
			return true
		}
	}
	return false
}

func RetrieveUserFromToken(token string) uint {
	_, _, claims := ValidateToken(token)
	return claims.UserID
}

func RemoveToken(token string) {
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, []byte(GetTokenSecret()))
	tokenBytes := []byte(token)
	var newClaims UserClaims
	jwt.ParseClaims(tokenBytes, verifier, &newClaims)
	DeleteToken(newClaims.ID)
}
