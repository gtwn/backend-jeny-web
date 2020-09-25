package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type User struct {
	ID          	primitive.ObjectID `bson:"_id,omitempty"`
	UserID      	string              `bson:"user_id,omitemty"`
	DisplayName 	[]string             `bson:"display,omitempty"`
}

type Task struct {
	ID          primitive.ObjectID 	`bson:"_id,omitempty"`
	OrderBy     string             	`bson:"order_by,omitempty"`
	Task        string             	`bson:"task,omitempty"`
	OrderTo     string             	`bson:"order_to,omitempty"`
	Deadline    time.Time          	`bson:"deadline,omitempty"`
	CreatedAt   time.Time          	`bson:"created_at,omitempty"`
	DoneAt      string          	`bson:"done_at,omitempty"`
	FromID		string				`bson:"from_id,omitempty"`
	GroupID		string				`bson:"group_id,omitempty"`
	Status 		string				`bson:"status,omitempty"`
}

type HistoryResponse struct {
	TaskHistory		[]Task
	FollowHistory	[]Task
}

type ReviewResponse struct {
	TaskReview		[]Task
	FollowReview	[]Task
}

type Msg struct {
	Type	string			`json:"type"`
	Text	string			`json:"text"`
}

type PushMsg struct {
	To 			string		`json:"to"`
	Message 	[]Msg		`json:"message"`
}