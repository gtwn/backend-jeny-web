package svc 

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByDisplay(Display string,userCollection *mongo.Collection) (*model.User,error) {

	var userResult model.User
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := userCollection.FindOne(ctx, bson.M{"display":Display}).Decode(&userResult) ;
	err != nil {
		return nil,err
	}

	return &userResult,nil
}