package mongo

import (
	"backendbase/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const userMongoCollection = "user"

var userCollection *mongo.Collection = GetCollection(Database, userMongoCollection)

func AllUser() ([]models.User, error) {
	var user models.User
	var users []models.User

	cursor, err := userCollection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		return users, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}

func FindUserByID(id string) (models.User, error) {
	var user models.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = userCollection.FindOne(Ctx, bson.D{{"_id", objectId}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SaveUser(user models.User) error {
	_, err := userCollection.InsertOne(Ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserByID(id string, user models.User) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
		"Age":       user.Age,
		"Gender":    user.Gender,
		"Email":     user.Email,
	}

	_, err = userCollection.UpdateOne(Ctx, bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserByID(id string) error {
	_, err := userCollection.DeleteOne(Ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
