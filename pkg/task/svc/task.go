package svc

import (
	"context"
	"time"

	"github.com/jenywebapp/pkg/task/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Task(userID string, taskCollection *mongo.Collection) (*[]model.Task,error) {
	var taskResult []model.Task

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	taskFind,err := taskCollection.Find(ctx, bson.M{"member_id": bson.M{"$elemMatch":bson.M{"$eq":userID}},"status":"In Progress"})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}

	if err := taskFind.All(ctx,&taskResult) ; err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}

	return &taskResult,nil
}