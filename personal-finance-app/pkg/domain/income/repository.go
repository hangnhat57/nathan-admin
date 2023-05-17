package income

import (
	"errors"
)

type IncomeRepository interface {
	Save(income *Income) error
	FindByID(id uint) (*Income, error)
	FindByUserID(userID uint) ([]Income, error)
	Update(income *Income) error
	Delete(id uint) error
}

var (
	ErrNoRecord = errors.New("no record found")
)
