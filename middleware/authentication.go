package middleware

import (
	"github.com/FelixWieland/eospke/config"
	"github.com/labstack/echo/v4/middleware"
)

// Authenticated checks if a user is authenticated
var Authenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(config.GithubClientSecret),
})
