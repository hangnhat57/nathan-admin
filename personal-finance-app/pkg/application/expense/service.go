package expense

import (
	"time"

	"personal-finance-app/pkg/domain/expense"
)

type Service struct {
	repo expense.ExpenseRepository
}

func NewExpenseService(repo expense.ExpenseRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddExpense(amount float64, description, category string, userID uint, date time.Time) error {
	// Here you would implement the logic for adding a new expense.
	// This could include validating the input, etc.

	newExpense := &expense.Expense{
		Amount:      amount,
		Description: description,
		Category:    category,
		UserId:      userID,
		Date:        date,
	}

	err := s.repo.Save(newExpense)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateExpense(id uint, amount float64, description, category string, date time.Time) error {
	// Here you would implement the logic for updating an existing expense.
	// This could include validating the input, checking if the expense with the given ID exists, etc.

	exp, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	exp.Amount = amount
	exp.Description = description
	exp.Category = category
	exp.Date = date

	err = s.repo.Update(exp)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteExpense(id uint) error {
	// Here you would implement the logic for deleting an expense.
	// This could include checking if the expense with the given ID exists, etc.

	return s.repo.Delete(id)
}

func (s *Service) GetExpenses(userID uint) ([]expense.Expense, error) {
	// Here you would implement the logic for retrieving all expenses of a user.

	return s.repo.FindByUserID(userID)
}
