package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Task(display []string, taskCollection *mongo.Collection) (*[]model.Task,error) {
	var taskResult []model.Task

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	for _,disp := range display {
		var task []model.Task
		taskFind,err := taskCollection.Find(ctx, bson.M{"order_to":disp})
		if err != nil {
			
		}
		if err := taskFind.All(ctx,&task); err == nil {
			taskResult = append(taskResult,task...)
		}
		
	}

	return &taskResult,nil
}