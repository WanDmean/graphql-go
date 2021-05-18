package todos

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoType struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text   string             `json:"text" bson:"text,omitempty"`
	Done   bool               `json:"done" bson:"done,omitempty"`
	UserID primitive.ObjectID `json:"userId" bson:"_user,omitempty"`
}
