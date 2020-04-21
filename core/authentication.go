package core

import (
	"time"

	"github.com/FelixWieland/eospke/config"
	"github.com/FelixWieland/eospke/types"
	"github.com/dgrijalva/jwt-go"
)

// Register -
func Register(username string, accessToken string) (types.Authenticated, error) {
	// add to database
	return Login(accessToken)
}

// Login -
func Login(accessToken string) (types.Authenticated, error) {
	// check if exists in database
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(config.SigningMethod, claims)
	stoken, err := token.SignedString([]byte(config.GithubClientSecret))
	if err != nil {
		return types.Authenticated{}, nil
	}
	return types.Authenticated{
		Bearer: stoken,
	}, nil
}
