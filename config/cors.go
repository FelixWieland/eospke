package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var CorsMiddleware = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"localhost" + Port},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
})
