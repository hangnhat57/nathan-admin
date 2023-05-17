package user

import (
	"errors"
)

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

var (
	ErrNoRecord       = errors.New("no record found")
	ErrDuplicateEmail = errors.New("duplicate email")
)
