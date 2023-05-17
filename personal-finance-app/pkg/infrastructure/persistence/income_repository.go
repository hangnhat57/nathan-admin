package persistence

import (
	"gorm.io/gorm"

	"personal-finance-app/pkg/domain/income"
)

type incomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) income.Repository {
	return &incomeRepository{
		db: db,
	}
}

func (r *incomeRepository) CreateIncome(income *income.Income) error {
	return r.db.Create(income).Error
}

func (r *incomeRepository) GetIncomes(userID uint) ([]*income.Income, error) {
	var incomes []*income.Income
	err := r.db.Where("user_id = ?", userID).Find(&incomes).Error
	if err != nil {
		return nil, err
	}

	return incomes, nil
}

func (r *incomeRepository) CreateIncomeSource(incomeSource *income.IncomeSource) error {
	return r.db.Create(incomeSource).Error
}

func (r *incomeRepository) GetIncomeSources(userID uint) ([]*income.IncomeSource, error) {
	var incomeSources []*income.IncomeSource
	err := r.db.Where("user_id = ?", userID).Find(&incomeSources).Error
	if err != nil {
		return nil, err
	}

	return incomeSources, nil
}
