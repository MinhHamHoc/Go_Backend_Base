package mongo

import (
	"backendbase/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const accountMongoCollection = "account"

var accountCollection *mongo.Collection = GetCollection(Database, accountMongoCollection)

func AllAccount() ([]models.Account, error) {
	var account models.Account
	var accounts []models.Account

	cursor, err := accountCollection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		return accounts, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&account)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, account)
	}

	return accounts, err
}

func FindAccountByID(id string) (models.Account, error) {
	var account models.Account
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, err
	}

	err = accountCollection.FindOne(Ctx, bson.D{{"_id", objectId}}).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func UpdateAccountByID(id string, account models.Account) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"Email":     account.Email,
		"Password":  account.Password,
		"UserID":    account.UserID,
		"CreatedBy": account.CreatedBy,
	}

	_, err = accountCollection.UpdateOne(Ctx, bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func SaveAccount(account models.Account) error {
	_, err := accountCollection.InsertOne(Ctx, account)
	if err != nil {
		return err
	}
	return nil
}

func RemoveAccountByID(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("invalid id")
	}

	_, err := accountCollection.DeleteOne(Ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
