package handler

import (
	"net/http"

	"github.com/FelixWieland/eospke/core"
	"github.com/FelixWieland/eospke/types"
	echo "github.com/labstack/echo/v4"
)

//AuthenticationHandler is a handler for User requests
type AuthenticationHandler Handler

//NewAuthenticationHandler creates a handler for authentication
func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}

// Register godoc
// @Summary Registers a User
// @Description Registers a User item
// @Tags Authorization
// @Accept json
// @Produce json
// @Param Authorization  body types.Authenticate true "New User"
// @Success 200 {object} types.Authenticated
// @Failure 404 {object} HTTPError
// @Router /authorization/register [post]
func (h *AuthenticationHandler) Register(c echo.Context) error {
	//username := ""
	//core.RegisterUser(username)
	return nil
}

// Login godoc
// @Summary Authenticates a User
// @Description Authenticates a User
// @Tags Authorization
// @Accept json
// @Produce json
// @Param Authorization body types.Authenticate true "Authorization"
// @Success 200 {object} types.Authenticated
// @Failure 404 {object} HTTPError
// @Router /authorization/login [post]
func (h *AuthenticationHandler) Login(c echo.Context) error {
	a := new(types.Authenticate)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(500, err)
	}
	loginResponse, err := core.Login(a.AccessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	return c.JSON(http.StatusOK, loginResponse)
}
