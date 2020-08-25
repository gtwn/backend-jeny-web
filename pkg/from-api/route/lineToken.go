package route

import (

	"github.com/davecgh/go-spew/spew"
	"github.com/jenywebapp/pkg/from-api/svc"
	"github.com/labstack/echo/v4"
)

type LineTokenConfig struct {
	LineAPI			string
	ChannelID		string
	ChannelSecret	string
}

func LineToken(cfg LineTokenConfig) echo.HandlerFunc {

	return func(c echo.Context) error {
		profile,err := svc.GetLineToken(cfg.LineAPI,cfg.ChannelID,cfg.ChannelSecret,c)
		if err != nil {
			return err
		}
		spew.Dump(profile)
		return c.JSON(200,profile)
	}
}