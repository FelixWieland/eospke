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

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://github.com/login/oauth/authorize?allow_signup=true&redirect_uri=https://localhost:1234/swagger/index.html

func main() {
	e := echo.New()
	e.Logger = echologrus.GetEchoLogger()

	e.Use(middleware.Cors)
	e.Use(middleware.Recover)

	e.HTTPErrorHandler = handler.ErrorHandler

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/request", config.RequestInformations)

	v1 := e.Group("/api/v1")
	v1.Use(middleware.Secure)
	v1.Use(middleware.Gzip)
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
			users.DELETE("/:id", h.DeleteUser)
			users.PATCH("/:id", h.UpdateUser)
		}
		authorization := v1.Group("/authorization")
		{
			h := handler.NewAuthenticationHandler()
			authorization.POST("/register", h.Register)
			authorization.POST("/login", h.Login)
		}

		v1.GET("/authenticated", func(c echo.Context) error {
			return c.String(200, "OK")
		}, middleware.Authenticated)

	}

	config.Serve(e)
}
