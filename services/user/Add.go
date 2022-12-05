package user

import (
	"backendbase/database/repository"
	"backendbase/models"
	utilities "backendbase/ultilities"
	"fmt"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"gopkg.in/mgo.v2/bson"
)

type AddUser struct {
	FirstName   string `bson:"first_name" json:"first_name"`
	LastName    string `bson:"last_name" json:"last_name"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Email       string `bson:"email" json:"email"`
	Gender      string `bson:"gender" json:"gender"`
	Age         int    `bson:"age" json:"age"`
}

func (a *AddUser) Valid() error {
	if a.PhoneNumber != "" && !utilities.IsPhoneNumber(a.PhoneNumber) {
		return fmt.Errorf("invalid phone number")
	}
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}

type AddUserHandler struct {
	UserRepository repository.UserRepository
}

func (h *AddUserHandler) Handle(a *AddUser) (string, error) {
	if err := a.Valid(); err != nil {
		return "", err
	}
	genderInt, err := strconv.Atoi(a.Gender)
	if err != nil || !models.IsValidGender(genderInt) {
		return "", err
	}
	u := models.User{
		ID:          bson.NewObjectId(),
		FirstName:   a.FirstName,
		LastName:    a.LastName,
		PhoneNumber: a.PhoneNumber,
		Gender:      genderInt,
		Age:         a.Age,
		CreatedTime: utilities.TimeInUTC(time.Now()),
		UpdatedTime: utilities.TimeInUTC(time.Now()),
	}
	return u.ID.Hex(), h.UserRepository.SaveUser(u)
}
