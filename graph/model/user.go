package model

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
