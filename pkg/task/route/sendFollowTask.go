package route

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"
	"github.com/jenywebapp/pkg/task/svc"
	"go.mongodb.org/mongo-driver/mongo"
)
type FollowTaskConfig struct {
	AccessToken	string
}

func SendFollowTask(cfg FollowTaskConfig,db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		id := c.Param("id")
		header := c.Request().Header.Get("Authorization")		// key IDToken
		payload,err := jwt.DecodeIDToken(header)
		expire := time.Unix(payload.Exp,0)
		user := db.Collection("user")
		taskCollection := db.Collection("task")
		if err != nil {
			return err
		}
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}
		// หางานด้วย ID
		task, err := svc.GetTaskByID(id,taskCollection)
		if err != nil {
			return err
		}
		// เอาชื่อ OrderTo หา User ในระบบ
		userResult, _ := svc.GetUserByDisplay(task.OrderTo,user) 
		
		if err := svc.PushMsgFollowTask(task,cfg.AccessToken,userResult,payload.Sub) ;
		err != nil {
			return err
		}

		return c.NoContent(200)
	}
}