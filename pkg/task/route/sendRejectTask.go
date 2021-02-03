package route

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"
	"github.com/jenywebapp/pkg/task/svc"
	"go.mongodb.org/mongo-driver/mongo"
)
type RejectTaskConfig struct {
	AccessToken	string
}

func SendRejectTask(cfg RejectTaskConfig,db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		id := c.Param("id")
		header := c.Request().Header.Get("Authorization")		// key IDToken
		if header == ""{
			return c.NoContent(401)
		}
		payload,err := jwt.DecodeIDToken(header)
		expire := time.Unix(payload.Exp,0)
		// user := db.Collection("user")
		taskCollection := db.Collection("tasks")
		if err != nil {
			return c.NoContent(400)
		}
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}
		// หางานด้วย ID
		task, err := svc.RejectTask(payload.Sub,id,taskCollection)
		if err != nil {
			return c.NoContent(400)
		}
		// เอาชื่อ OrderTo หา User ในระบบ
		// userResult, _ := svc.GetUserByDisplay(task.OrderTo,user) 
		
		if err := svc.PushMsgRejectTask(task,cfg.AccessToken) ;
		err != nil {
			return c.NoContent(400)
		}

		return c.NoContent(200)
	}
}