package route

import (
	"context"
	"fmt"
	"time"

	// "github.com/davecgh/go-spew/spew"

	"github.com/davecgh/go-spew/spew"
	"github.com/jenywebapp/pkg/from-api/model"
	"github.com/jenywebapp/pkg/from-api/svc"
	md "github.com/jenywebapp/pkg/jwt/model"
	"github.com/labstack/echo/v4"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	Task			[]model.Task
	FollowTask		[]model.Task
}

func LineToken(cfg LineTokenConfig,db *mongo.Database) echo.HandlerFunc {

	return func(c echo.Context) error {

		var userResult model.User
		var taskResult []model.Task
		var taskFollowResult []model.Task

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		userCollection := db.Collection("user")
		taskCollection := db.Collection("task")

		authSucess,payload,err := svc.GetLineToken(cfg.LineAPI,cfg.ChannelID,cfg.ChannelSecret,c)
		if err != nil {
			return err
		}
		// Display ไม่มี อาจจะเปลี่ยนชื่อหรือไม่มีตั้งแต่ต้น
		if err := userCollection.FindOne(ctx, bson.M{"display":payload.Name}).Decode(&userResult) ;err != nil {

			// Check ด้วย user_id ถ้ามี อัพเดท Display ไม่มี Insert ใหม่
			if err := userCollection.FindOne(ctx, bson.M{"user_id":payload.Sub}).Decode(&userResult) ; err != nil {
				if _, err := userCollection.InsertOne(ctx, model.User{
					UserID: payload.Sub,
					DisplayName: []string{payload.Name},
				}) ; err != nil {
					spew.Dump("Insert error",err)
					return err
				}
			} else {
				// อัพเดท display
				userResult.DisplayName = append(userResult.DisplayName,payload.Name)
				if _, err := userCollection.UpdateOne(
					ctx,
					bson.M{"user_id":payload.Sub},
					bson.D{
						{"$set", bson.M{"display":userResult.DisplayName}},
					}) ; err != nil {
						spew.Dump("Update error",err)
						return err
				}
			}
		}
		fmt.Print(userResult.DisplayName)
		taskFind,err := taskCollection.Find(ctx, bson.M{"order_to":payload.Name}) 
		if err != nil {
			// return err
		}
		if err := taskFind.All(ctx,&taskResult) ; err != nil {
			// return err
		}

		followTask,err := taskCollection.Find(ctx, bson.M{"order_by":payload.Name})
		if err != nil {

		}
		if err := followTask.All(ctx,&taskFollowResult) ; err != nil {

		}
		
		refresh := authSucess.RefreshToken
		expire := time.Unix(payload.Exp,0)

		return c.JSON(200,RespAuth{
			Payload: *payload,
			Refresh: refresh,
			Expire: expire,
			Task: taskResult,
			FollowTask: taskFollowResult,
		})
	}
}