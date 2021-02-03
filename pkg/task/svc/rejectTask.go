package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RejectTask(userID string,TaskID string,taskCollection *mongo.Collection) (*model.Task,error) {
	var stError error
	var taskResult model.Task
	id, _ := primitive.ObjectIDFromHex(TaskID)
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := taskCollection.FindOne(ctx, bson.M{"_id":id}).Decode(&taskResult) ;
	err != nil {
		return nil,err
	}
	if taskResult.FromID != userID {
		return nil, stError
	}
	
	if _,err := taskCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set" : bson.M{"status":"In Progress"}}) ; err != nil {
		return nil,err
	}
	
	

	return &taskResult,nil
}