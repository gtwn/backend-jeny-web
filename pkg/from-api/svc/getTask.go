package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/from-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetTask(userID string, taskCollection *mongo.Collection) (*[]model.Task,error) {
	var taskResult []model.Task

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	// for _,disp := range display {
	// 	var task []model.Task
	// 	taskFind,err := taskCollection.Find(ctx, bson.M{"order_id":userID,"status":"In Progress"})
	// 	if err != nil {
			
	// 	}
	// 	if err := taskFind.All(ctx,&task); err == nil {
	// 		taskResult = append(taskResult,task...)
	// 	}
		
	// }

	taskFind,err := taskCollection.Find(ctx, bson.M{"order_id":userID,"status":"In Progress"})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}
	if err := taskFind.All(ctx,&taskResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}

	return &taskResult,nil
}