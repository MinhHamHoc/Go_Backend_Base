package user

import (
	"backendbase/database/repository"
)

type DeleteUserHandler struct {
	UserRepository repository.UserRepository
}

func (h *DeleteUserHandler) Handle(id string) error {
	removeUser, err := h.UserRepository.FindByID(id)
	if err != nil {
		return err
	}

	err = h.UserRepository.RemoveByID(string(removeUser.ID))
	if err != nil {
		return err
	}
	return h.UserRepository.RemoveByID(removeUser.ID.Hex())
}
