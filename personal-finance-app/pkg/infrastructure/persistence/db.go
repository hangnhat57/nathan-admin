package persistence

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"personal-finance-app/pkg/domain/user"
	"personal-finance-app/pkg/domain/income"
	"personal-finance-app/pkg/domain/expense"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(user, password, host, name string, port int) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&user.User{}, &income.Income{}, &income.IncomeSource{}, &expense.Expense{})
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
