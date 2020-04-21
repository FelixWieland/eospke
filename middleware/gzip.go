package middleware

import (
	"compress/gzip"

	"github.com/labstack/echo/v4/middleware"
)

// Gzip compresses HTTP response using gzip compression scheme on the fly
var Gzip = middleware.GzipWithConfig(middleware.GzipConfig{
	Level: gzip.BestCompression,
})
