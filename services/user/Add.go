package user

import (
	"backendbase/database/repository"
	"backendbase/models"
	provider "backendbase/ultilities/providers"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gopkg.in/mgo.v2/bson"
)

type AddUser struct {
	FirstName   string `bson:"firstname"json:"firstname"`
	LastName    string `bson:"lastname"json:"lastname"`
	PhoneNumber string `bson:"phoneNumber"json:"phoneNumber"`
	Email       string `bson:"email"json:"email"`
	Gender      string `bson:"gender"json:"gender"`
}

func (a *AddUser) Valid() error {
	if a.PhoneNumber != "" && !provider.IsPhoneNumber(a.PhoneNumber) {
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
		// UserID:      a.ID.Hex(),
		// Age:         models.User,
		// CreatedTime: provider.TimeInUTC(time.Now()),
		// UpdatedTime: provider.TimeInUTC(time.Now()),
	}
	return u.ID.Hex(), h.UserRepository.Save(u)
}
