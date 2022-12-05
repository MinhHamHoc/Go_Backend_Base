package repository

import (
	"backendbase/models"
)

type UserRepository interface {
	AllUser() ([]models.User, error)
	FindUserByID(id string) (*models.User, error)
	SaveUser(user models.User) error
	UpdateUserByID(id string, user models.User) error
	RemoveUserByID(id string) error
}
