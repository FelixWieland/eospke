package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CorsMiddleware handles Cross-Origin-Resource-Sharing
var CorsMiddleware = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"localhost" + HTTPPort, "localhost" + HTTPSPort},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
})
