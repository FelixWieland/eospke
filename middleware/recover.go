package middleware

import "github.com/labstack/echo/v4/middleware"

// Recover recovers the server on error
var Recover = middleware.Recover()
