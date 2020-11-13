package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func History(userID string, taskCollection *mongo.Collection) (*[]model.Task,*[]model.Task,error) {
	var historyTaskResult []model.Task
	var historyFollowResult []model.Task


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	taskFind,err := taskCollection.Find(ctx, bson.M{"order_id":userID,"status":"Done"})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}
	if err := taskFind.All(ctx,&historyTaskResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}

	followFind,err := taskCollection.Find(ctx, bson.M{"from_id":userID,"status":"Done"}) 
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}

	if err := followFind.All(ctx,&historyFollowResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}


	return &historyTaskResult,&historyFollowResult,nil
}