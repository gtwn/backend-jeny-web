package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Follow(userID string, taskCollection *mongo.Collection) (*[]model.Task,error) {
	var followTaskResult []model.Task
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	// for _,disp := range display {
	// 	var task []model.Task
	// 	taskFind,err := taskCollection.Find(ctx, bson.M{"from_id":disp,"status":"In Progress"})
	// 	if err != nil {
			
	// 	}
	// 	if err := taskFind.All(ctx,&task); err == nil {
	// 		followTaskResult = append(followTaskResult,task...)
	// 	}
		
	// }
	taskFind,err := taskCollection.Find(ctx, bson.M{"from_id":userID,"status":"In Progress"})
	if err != nil && err != mongo.ErrNoDocuments{
		return nil,err
	}
	if err := taskFind.All(ctx,&followTaskResult); err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}

	

	return &followTaskResult,nil
}