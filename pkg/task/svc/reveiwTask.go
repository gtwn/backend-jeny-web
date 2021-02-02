package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReviewTask(TaskID string,taskCollection *mongo.Collection) (*model.Task,error) {

	var taskResult model.Task
	id, _ := primitive.ObjectIDFromHex(TaskID)
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := taskCollection.FindOne(ctx, bson.M{"_id":id}).Decode(&taskResult) ;
	err != nil {
		return nil,err
	}

	if taskResult.Type == "group" {
		_,err := taskCollection.UpdateMany(ctx, bson.M{"sub_id": taskResult.SubID}, bson.M{"$set" : bson.M{"status":"Review"}}) 
		if err != nil {
			return nil,err
		}
	} else {
		_,err := taskCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set" : bson.M{"status":"Review"}}) 
		if err != nil {
			return nil,err
		}
	}

	
	

	return &taskResult,nil
}