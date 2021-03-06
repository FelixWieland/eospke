package middleware

import "github.com/labstack/echo/v4/middleware"

// SecureMiddleware provides protection against cross-site scripting (XSS) attack, content type sniffing, clickjacking, insecure connection and other code injection attacks
var SecureMiddleware = middleware.SecureWithConfig(middleware.SecureConfig{
	XSSProtection:         "1; mode=block",
	ContentTypeNosniff:    "nosniff",
	XFrameOptions:         "SAMEORIGIN",
	HSTSMaxAge:            3600,
	ContentSecurityPolicy: "default-src 'self'",
})
