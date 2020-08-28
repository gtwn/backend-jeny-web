package route

import (
	"time"

	// "github.com/davecgh/go-spew/spew"

	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/jenywebapp/pkg/from-api/svc"
	md "github.com/jenywebapp/pkg/jwt/model"
	"github.com/labstack/echo/v4"
)

type LineTokenConfig struct {
	LineAPI			string
	ChannelID		string
	ChannelSecret	string
}

type RespAuth struct {
	Profile 		model.Profile
	Payload 		md.Payload
	Refresh			string
	Expire			time.Time
}

func LineToken(cfg LineTokenConfig) echo.HandlerFunc {

	return func(c echo.Context) error {
		// var cookie *http.Cookie
		authSucess,payload,err := svc.GetLineToken(cfg.LineAPI,cfg.ChannelID,cfg.ChannelSecret,c)
		if err != nil {
			return err
		}
		refresh := authSucess.RefreshToken
		expire := time.Unix(payload.Exp,0)

		return c.JSON(200,RespAuth{
			Payload: *payload,
			Refresh: refresh,
			Expire: expire,
		})
	}
}