package route

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jenywebapp/pkg/jwt"
	sv "github.com/jenywebapp/pkg/svc"
	"github.com/jenywebapp/pkg/task/model"
	"github.com/jenywebapp/pkg/task/svc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetReview(db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {
		
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		userCollection := db.Collection("user")
		taskCollection := db.Collection("task")
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
		var user model.User
		if err := userCollection.FindOne(ctx, bson.M{"user_id":payload.Sub}).Decode(&user) ; err != nil {
			return err
		}

		reviewTask,reviewFollow,err := svc.Review(user.DisplayName,taskCollection)
		if err != nil {
			return err
		}
		
		return c.JSON(200,model.ReviewResponse{
			TaskReview: *reviewTask,
			FollowReview: *reviewFollow,
		})
	}
}