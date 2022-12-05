package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password,omitempty" json:"-"`
	UserID    string        `bson:"userID" json:"userID"`
	CreatedBy string        `bson:"createdBy" json:"createdBy"`
}
