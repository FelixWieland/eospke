package middleware

import (
	"github.com/FelixWieland/eospke/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CorsMiddleware handles Cross-Origin-Resource-Sharing
var CorsMiddleware = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"localhost" + config.HTTPPort, "localhost" + config.HTTPSPort},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
})
