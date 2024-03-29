package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"user_id,omitemty"`
	DisplayName []string           `bson:"display,omitempty"`
}

// type Task struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	OrderBy   string             `bson:"order_by,omitempty"`
// 	Task      string             `bson:"task,omitempty"`
// 	OrderTo   string             `bson:"order_to,omitempty"`
// 	Deadline  time.Time          `bson:"deadline,omitempty"`
// 	CreatedAt time.Time          `bson:"created_at,omitempty"`
// 	DoneAt    string             `bson:"done_at,omitempty"`
// 	FromID    string             `bson:"from_id,omitempty"`
// 	OrderID	  string			 `bson:"order_id,omitempty"`
// 	GroupID   string             `bson:"group_id,omitempty"`
// 	Status    string             `bson:"status,omitempty"`
// }

type Task struct {
	ID          primitive.ObjectID 	`bson:"_id,omitempty" json:"id"` 
	SubID		string				`bson:"sub_id" json:"sub_id"`
	OrderBy     string             	`bson:"order_by,omitempty" json:"order_by"`
	Task        string             	`bson:"task,omitempty" json:"task"`
	Detail		string				`bson:"detail,omitempty" json:"detail"`
	OrderTo     string             	`bson:"order_to,omitempty" json:"order_to"`
	Deadline    time.Time          	`bson:"deadline,omitempty" json:"deadline"`
	CreatedAt   time.Time          	`bson:"created_at,omitempty" json:"created_at"`
	DoneAt      time.Time          	`bson:"done_at,omitempty" json:"done_at"`
	FromID		string				`bson:"from_id,omitempty" json:"from_id"`
	OrderID		string				`bson:"order_id,omitempty" json:"order_id"`
	GroupID		string				`bson:"group_id,omitempty" json:"group_id"`
	Member		[]string			`bson:"member,omitempty" json:"member"`
	MemberID    []string			`bson:"member_id,omitempty" json:"member_id"`
	Type 		string				`bson:"type,omitempty"	json:"tpye"`
	Status 		string				`bson:"status,omitempty" json:"status"`
}