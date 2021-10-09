package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id"   bson: "id"`
	Name     string        `json:"name" bson: "name"`
	Email    string        `json:"email" bson: "email"`
	Password int           `json:"password" bson: "password"`
}
