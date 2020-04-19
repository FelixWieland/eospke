package handler

import (
	"github.com/labstack/echo/v4"
)

//SurveyHandler is a handler for Survey requests
type SurveyHandler Handler

//NewSurveyHandler creates a handler for Survey requests
func NewSurveyHandler() *SurveyHandler {
	return &SurveyHandler{}
}

// GetSurveys godoc
// @Summary Returns all Surveys
// @Description Returns all survey items
// @Tags Survey
// @Accept json
// @Produce json
// @Success 200 {object} types.Surveys
// @Failure 400 {object} HTTPError
// @Router /surveys [get]
func (h *SurveyHandler) GetSurveys(c echo.Context) error {
	return nil
}

// FindSurvey godoc
// @Summary Finds a Survey
// @Description Returns a Survey item
// @Tags Survey
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} types.Survey
// @Failure 400 {object} HTTPError
// @Router /surveys/{id} [get]
func (h *SurveyHandler) FindSurvey(c echo.Context) error {
	return nil
}

// CreateSurvey godoc
// @Summary Create a Survey
// @Description Creates a Survey item
// @Tags Survey
// @Accept json
// @Produce json
// @Param todo body types.Survey true "New Survey"
// @Success 201 {object} types.Survey
// @Failure 400 {object} HTTPError
// @Router /surveys [post]
func (h *SurveyHandler) CreateSurvey(c echo.Context) error {
	return nil
}

// UpdateSurvey godoc
// @Summary Update a Survey
// @Description Updates a Survey item
// @Tags Survey
// @Accept json
// @Produce json
// @Param todo body types.Survey true "Updated Survey"
// @Param id path int true "Survey ID"
// @Success 200 {object} types.Survey
// @Failure 404 {object} HTTPError
// @Router /surveys/{id} [patch]
func (h *SurveyHandler) UpdateSurvey(c echo.Context) error {
	return nil
}

// DeleteSurvey godoc
// @Summary Delete a Survey
// @Description Deletes a Survey item
// @Tags Survey
// @Produce json
// @Param id path int true "Survey ID"
// @Success 204 {object} types.Survey
// @Failure 404 {object} HTTPError
// @Router /surveys/{id} [delete]
func (h *SurveyHandler) DeleteSurvey(c echo.Context) error {
	return nil
}
