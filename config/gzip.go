package config

import (
	"compress/gzip"

	"github.com/labstack/echo/v4/middleware"
)

// GzipMiddleware compresses HTTP response using gzip compression scheme on the fly
var GzipMiddleware = middleware.GzipWithConfig(middleware.GzipConfig{
	Level: gzip.BestCompression,
})
