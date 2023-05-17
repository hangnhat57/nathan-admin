package persistence

import (
	"fmt"
	"gorm.io/gorm"
	"personal-finance-app/pkg/domain/user"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *user.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *user.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) FindByID(id uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("not found")
	}
	return &user, nil
}
