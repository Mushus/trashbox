package user

import "github.com/Mushus/trashbox/backend/server/app/property"

// Repository
type Repository interface {
	FindByLogin(login string) (*property.User, error)
	Add(user *property.User) error
}
