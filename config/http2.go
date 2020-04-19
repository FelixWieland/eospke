package config

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
