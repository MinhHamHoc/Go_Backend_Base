package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Age       int           `bson:"age" json:"age"`
	Gender    int           `bson:"gender" json:"gender"`
	Email     string        `bson:"email" json:"email"`
}
