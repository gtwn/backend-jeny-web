package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type User struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty"`
	UserID		int64				`bson:"user_id,omitemty"`
	DisplayName	string				`bson:"display,omitempty"`
	ProfileURL	string				`bson:"url,omitempty"`
}

type Task	struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty"`
	OrderBy		string				`bson:"order_by,omitempty"`
	Task		string				`bson:"task,omitempty"`
	OrderTo		string				`bson:"order_to,omitempty"`
	Deadline	time.Time			`bson:"deadline,omitempty"`
	CreatedAt	time.Time			`bson:"created_at,omitempty"`
	DoneAt		time.Time			`bson:"done_at,omitempty"`
	GroupID		int64				`bson:"group_id,omitempty"`
	OrderUserID	int64				`bson:"order_user_id,omitempty"`
	Status		string				`bson:"status,omitempty"`

}