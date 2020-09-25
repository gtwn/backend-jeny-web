package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskByID(TaskID string,taskCollection *mongo.Collection) (*model.Task,error) {

	var taskResult model.Task

	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := taskCollection.FindOne(ctx, bson.M{"_id":TaskID}).Decode(&taskResult) ;
	err != nil {
		return nil,err
	}

	_,err := taskCollection.UpdateOne(ctx, bson.M{"_id": TaskID}, bson.M{"$set" : bson.M{"status":"Review"}}) 
	if err != nil {
		return nil,err
	}
	

	return &taskResult,nil
}