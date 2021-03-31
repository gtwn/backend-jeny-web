package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type User struct {
	ID          	primitive.ObjectID 	`bson:"_id,omitempty"`
	UserID      	string              `bson:"user_id,omitemty"`
	DisplayName 	[]string            `bson:"display,omitempty"`
}

type Task struct {
	ID          primitive.ObjectID 	`bson:"_id,omitempty" json:"id"` 
	OrderBy     string             	`bson:"order_by,omitempty" json:"order_by"`
	Task        string             	`bson:"task,omitempty" json:"task"`
	Detail		string				`bson:"detail,omitempty" json:"detail"`
	Deadline    time.Time          	`bson:"deadline,omitempty" json:"deadline"`
	CreatedAt   time.Time          	`bson:"created_at,omitempty" json:"created_at"`
	DoneAt      time.Time          	`bson:"done_at,omitempty" json:"done_at"`
	FromID		string				`bson:"from_id,omitempty" json:"from_id"`
	GroupID		string				`bson:"group_id,omitempty" json:"group_id"`
	GroupName 	string				`bson:"group_name",omitempty json:"group_name"`
	Member		[]string			`bson:"member,omitempty" json:"member"`
	MemberID    []string			`bson:"member_id,omitempty" json:"member_id"`
	Type 		string				`bson:"type,omitempty"	json:"tpye"`
	Status 		string				`bson:"status,omitempty" json:"status"`
}

type TaskResponse struct {
	Task 		[]Task			`json:"task"`
}

type HistoryResponse struct {
	TaskHistory		[]Task		`json:"task_history"`
	FollowHistory	[]Task		`json:"follow_history"`
}

type ReviewResponse struct {
	TaskReview		[]Task		`json:"task_review"`
	FollowReview	[]Task		`json:"follow_review"`
}

type FollowResponse struct {
	Follow 			[]Task		`json:"follow"`
}

type Msg struct {
	Type	string				`json:"type"`
	Text	string				`json:"text"`
}

type PushMsg struct {
	To 			string			`json:"to"`
	Message 	[]Msg			`json:"messages"`
}


type PushMultiple struct {
	To			[]string		`json:"to"`
	Message 	[]Msg			`json:"messages"`
}