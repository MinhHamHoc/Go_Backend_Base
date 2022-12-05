package repository

import (
	"backendbase/models"
)

type AccountRepository interface {
	AllAccount() ([]models.User, error)
	FindAccountByIDByID(id string) (*models.User, error)
	SaveAccount(user models.User) error
	UpdateAccountByID(id string, user models.User) error
	RemoveAccountByIDByID(id string) error
}
