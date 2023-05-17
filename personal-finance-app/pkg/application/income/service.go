package income

import (
	"errors"
	"time"

	"personal-finance-app/pkg/domain/income"
)

type IncomeService struct {
	repo income.IncomeRepository
}

func NewIncomeService(repo income.IncomeRepository) *IncomeService {
	return &IncomeService{repo: repo}
}

func (s *IncomeService) AddIncome(amount float64, description, source string, userID uint, date time.Time) error {
	// Here you would implement the logic for adding a new income.
	// This could include validating the input, etc.

	newIncome := &income.Income{
		Amount:      amount,
		Description: description,
		Source:      source,
		UserId:      userID,
		Date:        date,
	}

	err := s.repo.Save(newIncome)
	if err != nil {
		return err
	}

	return nil
}

func (s *IncomeService) UpdateIncome(id uint, amount float64, description, source string, date time.Time) error {
	// Here you would implement the logic for updating an existing income.
	// This could include validating the input, checking if the income with the given ID exists, etc.

	inc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	inc.Amount = amount
	inc.Description = description
	inc.Source = source
	inc.Date = date

	err = s.repo.Update(inc)
	if err != nil {
		return err
	}

	return nil
}

func (s *IncomeService) DeleteIncome(id uint) error {
	// Here you would implement the logic for deleting an income.
	// This could include checking if the income with the given ID exists, etc.

	return s.repo.Delete(id)
}

func (s *IncomeService) GetIncomes(userID uint) ([]income.Income, error) {
	// Here you would implement the logic for retrieving all incomes of a user.

	return s.repo.FindByUserID(userID)
}
