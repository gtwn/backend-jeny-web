package route

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"
	"github.com/jenywebapp/pkg/task/svc"
	"go.mongodb.org/mongo-driver/mongo"
)
type cfgReviewTask struct {
	AccessToken	string
}

func SendReviewTask(db *mongo.Database,cfg cfgReviewTask) echo.HandlerFunc {

	return func(c echo.Context) error {
		id := c.Param("id")
		header := c.Request().Header.Get("Authorization")		// key IDToken
		payload,err := jwt.DecodeIDToken(header)
		expire := time.Unix(payload.Exp,0)
		taskCollection := db.Collection("task")
		if err != nil {
			return err
		}
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}

		task, err := svc.GetTaskByID(id,taskCollection)
		if err != nil {
			return err
		}
		if err := svc.PushMsgSendTask(task,cfg.AccessToken,payload.Name,payload.Sub) ;
		err != nil {
			return err
		}

		return c.NoContent(200)
	}
}