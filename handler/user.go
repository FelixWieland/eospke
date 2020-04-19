package handler

import echo "github.com/labstack/echo/v4"

//UserHandler is a handler for User requests
type UserHandler Handler

//NewUserHandler creates a handler for User requests
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUsers godoc
// @Summary Returns all Users
// @Description Returns all User items
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} types.Users
// @Failure 404 {object} HTTPError
// @Router /users [get]
func (h *UserHandler) GetUsers(c echo.Context) error {
	return nil
}

// FindUser godoc
// @Summary Returns a User
// @Description Returns a User item
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} types.User
// @Failure 404 {object} HTTPError
// @Router /users/{id} [get]
func (h *UserHandler) FindUser(c echo.Context) error {
	return nil
}

// RegisterUser godoc
// @Summary Creates a User
// @Description Creates a User item
// @Tags User
// @Accept json
// @Produce json
// @Param todo body types.User true "New User"
// @Success 200 {object} types.User
// @Failure 404 {object} HTTPError
// @Router /users [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	return nil
}

// UpdateUser godoc
// @Summary Updates a User
// @Description Updates a User item
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param todo body types.User true "Updated User"
// @Success 200 {object} types.User
// @Failure 404 {object} HTTPError
// @Router /users/{id} [patch]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser godoc
// @Summary Delete a User
// @Description Deletes a User item
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} types.User
// @Failure 404 {object} HTTPError
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	return nil
}
