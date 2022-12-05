package models

type Book struct {
	Name   string `bson:"name" json:"name"`
	Author string `bson:"author" json:"author"`
	
}
