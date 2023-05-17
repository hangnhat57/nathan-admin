package user

import (
	"errors"
)

type Repository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
}

var (
	ErrNoRecord       = errors.New("no record found")
	ErrDuplicateEmail = errors.New("duplicate email")
)
