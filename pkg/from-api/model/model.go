package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type AuthSuccess struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	IDToken      string  `json:"id_token"`
	RefreshToken string  `json:"refresh_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}

type Profile struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string              `bson:"user_id,omitemty"`
	DisplayName []string             `bson:"display,omitempty"`
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