package expense

import (
	"errors"
)

type ExpenseRepository interface {
	Save(expense *Expense) error
	FindByID(id uint) (*Expense, error)
	FindByUserID(userID uint) ([]Expense, error)
	Update(expense *Expense) error
	Delete(id uint) error
}

var (
	ErrNoRecord = errors.New("no record found")
)
