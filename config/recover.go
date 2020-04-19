package config

import "github.com/labstack/echo/v4/middleware"

// RecoverMiddleware recovers the server on error
var RecoverMiddleware = middleware.Recover()
