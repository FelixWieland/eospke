package middleware

import (
	"github.com/FelixWieland/eospke/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors handles Cross-Origin-Resource-Sharing
var Cors = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"localhost" + config.HTTPPort, "localhost" + config.HTTPSPort},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
})
