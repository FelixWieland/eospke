package config

import "github.com/labstack/echo/v4"

// HTTPSPort is the Port for a HTTPS Connection
const HTTPSPort = ":443"

// HTTPPort is the Port for a HTTP Connection
const HTTPPort = ":80"

// Serve starts the echo server
func Serve(e *echo.Echo) {
	go func() {
		e.Logger.Fatal(e.Start(HTTPPort))
	}()
	e.Logger.Fatal(e.StartTLS(HTTPSPort, "cert.pem", "key.pem"))
}
