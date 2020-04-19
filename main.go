package main

import (
	"github.com/FelixWieland/eospke/config"
	_ "github.com/FelixWieland/eospke/docs"
	"github.com/FelixWieland/eospke/handler"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Survey API
// @version 1.0
// @description This is a sample how to use Echo to create restful APIs

// @contact.name Felix Wieland
// @contact.email felix.wieland@mail.de

// @BasePath /api/v1
func main() {
	e := echo.New()

	e.Use(config.CorsMiddleware)
	e.Use(config.SecureMiddleware)
	e.Use(config.GzipMiddleware)
	e.Use(config.RecoverMiddleware)

	e.HTTPErrorHandler = handler.ErrorHandler

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/request", config.RequestInformations)

	v1 := e.Group("/api/v1")
	{
		surveys := v1.Group("/surveys")
		{
			h := handler.NewSurveyHandler()
			surveys.GET("/", h.GetSurveys)
			surveys.GET("/:id", h.FindSurvey)
			surveys.POST("/", h.CreateSurvey)
			surveys.PATCH("/:id", h.UpdateSurvey)
			surveys.DELETE("/:id", h.DeleteSurvey)
		}
		users := v1.Group("/users")
		{
			h := handler.NewUserHandler()
			users.GET("/", h.GetUsers)
			users.GET("/:id", h.FindUser)
			users.POST("/", h.RegisterUser)
			users.DELETE("/:id", h.DeleteUser)
			users.PATCH("/:id", h.UpdateUser)
		}
	}

	config.Serve(e)
}
