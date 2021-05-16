package model

type Todo struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Text   string `json:"text" bson:"text,omitempty"`
	Done   bool   `json:"done" bson:"done,omitempty"`
	UserID string `json:"userId" bson:"_user,omitempty"`
}

type NewTodo struct {
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}
