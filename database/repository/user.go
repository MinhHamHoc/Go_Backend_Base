package repository

import (
	"backendbase/models"
)

type UserRepository interface {
	All() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	Save(user models.User) error
	UpdateByID(id string, user models.User) error
	RemoveByID(id string) error
}
