package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Review(display []string, taskCollection *mongo.Collection) (*[]model.Task,*[]model.Task,error) {
	var ReviewTaskResult []model.Task
	var ReviewFollowResult []model.Task


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	for _,disp := range display {
		var task []model.Task
		var follow []model.Task
		taskFind,err := taskCollection.Find(ctx, bson.M{"order_to":disp,"status":"Review"})
		if err != nil {
			
		}
		if err := taskFind.All(ctx,&task); err == nil {
			ReviewTaskResult = append(ReviewTaskResult,task...)
		}

		followFind,err := taskCollection.Find(ctx, bson.M{"order_by":disp,"status":"Review"})
		if err != nil {
			
		}
		if err := followFind.All(ctx,&follow) ; err == nil {
			ReviewFollowResult = append(ReviewFollowResult,follow...)
		}
		
	}


	return &ReviewTaskResult,&ReviewFollowResult,nil
}