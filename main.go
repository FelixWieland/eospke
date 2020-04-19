package main

import (
	"github.com/FelixWieland/eospke/config"
	_ "github.com/FelixWieland/eospke/docs"
	"github.com/FelixWieland/eospke/handler"
	"github.com/FelixWieland/eospke/middleware"
	echologrus "github.com/plutov/echo-logrus"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	echologrus.Logger = logrus.New()
	elasticHook, err := config.MakeElasticHook(echologrus.Logger)
	if err != nil {
		echologrus.Logger.Errorf("Could not connect to elasticsearch: %v", err)
		return
	}
	echologrus.Logger.AddHook(elasticHook)
	echologrus.Logger.Info("Initialized logrus")
}

// @title Echo Survey API
// @version 1.0
// @description This is a sample how to use Echo to create restful APIs

// @contact.name Felix Wieland
// @contact.email felix.wieland@mail.de

// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Logger = echologrus.GetEchoLogger()

	e.Use(middleware.CorsMiddleware)
	e.Use(middleware.SecureMiddleware)
	e.Use(middleware.GzipMiddleware)
	e.Use(middleware.RecoverMiddleware)

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
