package route

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"
	"github.com/jenywebapp/pkg/task/model"
	"github.com/jenywebapp/pkg/task/svc"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetHistory(db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		
		taskCollection := db.Collection("tasks")
		header := c.Request().Header.Get("Authorization")		// key IDToken
		if header == ""{
			return c.NoContent(401)
		}
		payload,err := jwt.DecodeIDToken(header)
		expire := time.Unix(payload.Exp,0)
		if err != nil {
			return err
		}
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}
	

		historyTask,historyFollow,err := svc.History(payload.Sub,taskCollection)
		if err != nil {
			return c.NoContent(400)
		}
		
		return c.JSON(200,model.HistoryResponse{
			TaskHistory: *historyTask,
			FollowHistory: *historyFollow,
		})
	}
}