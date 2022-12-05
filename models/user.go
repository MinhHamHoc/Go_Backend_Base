package models

import (
	"go.starlark.net/lib/time"
	"gopkg.in/mgo.v2/bson"
)

const (
	Male   = 1
	FeMale = 0
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	FirstName   string        `bson:"firstName" json:"firstName"`
	LastName    string        `bson:"lastName" json:"lastName"`
	Age         int           `bson:"age" json:"age"`
	Gender      int           `bson:"gender" json:"gender"`
	PhoneNumber string        `bson:"phoneNumber"json:"phoneNumber"`
	Email       string        `bson:"email" json:"email"`
	UserID      string        `bson:"userID" json:"userID"`
	CreatedTime time.Time     `bson :"createdTime"json:"createdTime"`
	UpdatedTime time.Time     `bson:"updatedTime"json:"updatedTime"`
}

func IsValidGender(genderInt int) bool {
	if genderInt != Male && genderInt != FeMale {
		return false
	}
	return true
}
