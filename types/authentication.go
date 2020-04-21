package types

import "github.com/dgrijalva/jwt-go"

// Authenticate handels logins and registers
type Authenticate struct {
	AccessToken string
}

// EopkeClaims extends the default jwt claims
type EopkeClaims struct {
	jwt.StandardClaims
}

// Authenticated contains the jwt Bearer
type Authenticated struct {
	Bearer string
}
