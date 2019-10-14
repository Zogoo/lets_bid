package models

type User struct {
	id       string `json:"user_id" bson:"userId"`
	name     string `json:"name" bson:"Name"`
	email    string `json:"email" bson:"Email"`
	password string `json:"password" bson:"Password"`
}
