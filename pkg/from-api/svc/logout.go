package svc

import (

	resty "github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func LogOut(c echo.Context) error {

	client := resty.New()
	resfresh := c.Request().Header.Get("resfresh_token")
	if _,err := client.R().
		SetHeader("Content-Type","application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"refresh_token": resfresh,
		}).Post("https://api.line.me/v2/oauth/revoke") ; err != nil {
			return err
		}
	return nil
}