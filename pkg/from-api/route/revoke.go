package route

import (
	"github.com/jenywebapp/pkg/from-api/svc"
	"github.com/labstack/echo/v4"
)

func Revoke() echo.HandlerFunc {

	return func(c echo.Context) error {
		if err := svc.LogOut(c) ; err != nil {
			return err
		}
		return c.NoContent(200)
	}
}