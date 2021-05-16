package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text   string             `json:"text" bson:"text,omitempty"`
	Done   bool               `json:"done" bson:"done,omitempty"`
	UserID primitive.ObjectID `json:"user" bson:"user,omitempty"`
}

type NewTodo struct {
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}
