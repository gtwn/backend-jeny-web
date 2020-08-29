package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func History(display []string, taskCollection *mongo.Collection) (*[]model.Task,*[]model.Task,error) {
	var historyTaskResult []model.Task
	var historyFollowResult []model.Task


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	for _,disp := range display {
		var task []model.Task
		var follow []model.Task
		taskFind,err := taskCollection.Find(ctx, bson.M{"order_to":disp,"status":"Done"})
		if err != nil {
			
		}
		if err := taskFind.All(ctx,&task); err == nil {
			historyTaskResult = append(historyTaskResult,task...)
		}

		followFind,err := taskCollection.Find(ctx, bson.M{"order_by":disp,"status":"Done"})
		if err != nil {
			
		}
		if err := followFind.All(ctx,&follow) ; err == nil {
			historyFollowResult = append(historyFollowResult,follow...)
		}
		
	}


	return &historyTaskResult,&historyFollowResult,nil
}