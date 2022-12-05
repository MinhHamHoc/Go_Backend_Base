package user

import (
	"backendbase/database/repository"
)

type UpdateUser struct {
	FirstName   string `bson:"firstname"json:"firstname"`
	LastName    string `bson:"lastname"json:"lastname"`
	PhoneNumber string `bson:"phoneNumber"json:"phoneNumber"`
	Email       string `bson:"email"json:"email"`
}

func (c *UpdateUser) Valid() error {
	return nil
}

type UpdateUserHandler struct {
	UserRepository repository.UserRepository
}

func (h *UpdateUserHandler) Handle(a *UpdateUser) error {
	err := a.Valid()
	if err != nil {
		return err
	}
	userUpdate, err := h.UserRepository.FindUserByID(a.PhoneNumber)
	if err != nil {
		return err
	}
	if a.FirstName != "" {
		userUpdate.FirstName = a.FirstName
	}
	if a.LastName != "" {
		userUpdate.LastName = a.LastName
	}
	if a.PhoneNumber != "" {
		userUpdate.PhoneNumber = a.PhoneNumber
	}
	// userUpdate.UpdatedTime = provider.TimeInUTC(time.Now())
	return h.UserRepository.UpdateUserByID(a.PhoneNumber, *userUpdate)
}
