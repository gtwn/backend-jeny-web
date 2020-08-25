package main

import (
	"fmt"
	"net/http"

	"github.com/jenywebapp/cmd/jeny/config"
	route "github.com/jenywebapp/pkg/from-api/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg,err := config.Read()
	if err != nil {
		logrus.WithError(err).Fatal("Read config")
	}
	
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Middleware debug
	e.Use(middleware.Logger())

	e.GET("/callback", route.LineToken(route.LineTokenConfig{
		LineAPI 		: cfg.LineTokenAPI,
		ChannelID 		: cfg.ChannelID,
		ChannelSecret	: cfg.ChannelSecret,
	}))

	e.Logger.Fatal(e.Start(fmt.Sprint(":", cfg.Port)))
}