package route

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/jenywebapp/pkg/from-api/svc"
	"github.com/labstack/echo/v4"
)

type LineTokenConfig struct {
	LineAPI			string
	ChannelID		string
	ChannelSecret	string
}

type RespAuth struct {
	Profile 		model.Profile
	Auth 			model.AuthSuccess
}

func LineToken(cfg LineTokenConfig) echo.HandlerFunc {

	return func(c echo.Context) error {
		authSucess,profile,err := svc.GetLineToken(cfg.LineAPI,cfg.ChannelID,cfg.ChannelSecret,c)
		if err != nil {
			return err
		}
		spew.Dump(profile,"\n",authSucess)
		return c.JSON(200,RespAuth{
			Profile: *profile,
			Auth: *authSucess,
		})
	}
}