package persistence

import (
	"gorm.io/gorm"

	"personal-finance-app/pkg/domain/expense"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) expense.Repository {
	return &expenseRepository{
		db: db,
	}
}

func (r *expenseRepository) CreateExpense(expense *expense.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) GetExpenses(userID uint) ([]*expense.Expense, error) {
	var expenses []*expense.Expense
	err := r.db.Where("user_id = ?", userID).Find(&expenses).Error
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
