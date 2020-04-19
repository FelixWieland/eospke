package config

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

//RequestInformations returns informations about the client request
func RequestInformations(c echo.Context) error {
	req := c.Request()
	format := `
	  <code>
		Protocol: %s<br>
		Host: %s<br>
		Remote Address: %s<br>
		Method: %s<br>
		Path: %s<br>
	  </code>
	`
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}
