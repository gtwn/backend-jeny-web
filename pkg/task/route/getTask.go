package route

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetTask(db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		var taskResult []model.Task
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		taskCollection := db.Collection("task")
		name := c.QueryParam("name")
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return c.NoContent(http.StatusUnauthorized)
		}
		taskFind,err := taskCollection.Find(ctx, bson.M{"order_to":name}) 
		if err != nil {
			// return err
		}
		if err := taskFind.All(ctx,&taskResult) ; err != nil {
			// return err
		}
		return c.JSON(200,taskResult)
	}
}