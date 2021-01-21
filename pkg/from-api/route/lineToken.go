package route

import (
	"time"

	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/jenywebapp/pkg/from-api/svc"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
)

type LineTokenConfig struct {
	LineAPI			string
	ChannelID		string
	ChannelSecret	string
}

type RespAuth struct {

	Task			[]model.Task		`json:"task"`
	FollowTask		[]model.Task		`json:"follow"`
	Token			string				`json:"token"`
}

func LineToken(cfg LineTokenConfig,db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {

		taskCollection := db.Collection("tasklist")

		// authSucess,payload,err := svc.GetLineToken(cfg.LineAPI,cfg.ChannelID,cfg.ChannelSecret,c)
		// if err != nil {
		// 	return err
		// }
		// refresh := authSucess.RefreshToken
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return c.NoContent(401)
		}
		payload, err := jwt.DecodeIDToken(header)
		if err != nil {
			return err
		}
		expire := time.Unix(payload.Exp,0)
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}


		taskResp,err := svc.GetTask(payload.Sub,taskCollection)
		followResp,err := svc.GetFollow(payload.Sub,taskCollection)
		

		return c.JSON(200,RespAuth{
			Task: *taskResp,
			FollowTask: *followResp,
			// Token: authSucess.IDToken,
		})
	}
}