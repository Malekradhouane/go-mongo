package store

import (
	"github/malekradhouane/test-cdi/api"
)

type Store interface {
	UserStore
}

//UserStore represents the interface to manage users storage
type UserStore interface {
	CreateUser(*User) (*User, error)
	IsEmailTaken(string) bool
	Authenticate(*api.Login) (*User, error)
	GetAllUsers() ([]*User, error)
	Get(string) (*User, error)
	DeleteUser(string) error
	UpdateUser(*User, string)  error
}
