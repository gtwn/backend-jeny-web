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


func GetTask(db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		
		taskCollection := db.Collection("task")
		header := c.Request().Header.Get("Authorization")		// key IDToken
		if header == ""{
			return c.NoContent(401)
		}
		payload,err := jwt.DecodeIDToken(header)
		expire := time.Unix(payload.Exp,0)
		if err != nil {
			return c.NoContent(400)
		}
		if sv.CheckExpire(expire) != true {
			return c.NoContent(401)
		}
	
		task,err := svc.Task(payload.Sub,taskCollection)
		if err != nil {
			return c.NoContent(400)
		}
		
		return c.JSON(200,model.TaskResponse{
			Task: *task,
		})
	}
}