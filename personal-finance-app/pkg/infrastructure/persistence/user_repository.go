package persistence

import (
	"gorm.io/gorm"

	"personal-finance-app/pkg/domain/user"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(userID uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(user *user.User) error {
	return r.db.Save(user).Error
}
