package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Review(userID string, taskCollection *mongo.Collection) (*[]model.Task,*[]model.Task,error) {
	var ReviewTaskResult []model.Task
	var ReviewFollowResult []model.Task


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	
	taskFind,err := taskCollection.Find(ctx, bson.M{"member_id": bson.M{"$elemMatch":bson.M{"$eq":userID}},"status":"Review"})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}

	if err := taskFind.All(ctx,&ReviewTaskResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}

	followFind, err := taskCollection.Find(ctx,bson.M{"from_id":userID,"status":"Review"})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}
	
	if err := followFind.All(ctx,&ReviewFollowResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,nil,err
	}

	return &ReviewTaskResult,&ReviewFollowResult,nil
}