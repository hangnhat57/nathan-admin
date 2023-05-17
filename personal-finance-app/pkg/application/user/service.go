package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"personal-finance-app/pkg/domain/user"
)

type UserService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(email, password string) (*user.User, error) {
	// Here you would implement the logic for user registration.
	// This could include validating the input, checking if a user with the given email already exists, etc.
	// Then, hash the password before storing it in the database.

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.Save(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	// Here you would implement the logic for user login.
	// This could include validating the input, checking if a user with the given email exists, etc.
	// Then, compare the given password with the hashed password in the database.

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Then, generate and return a JWT token.

	return "", nil
}

func (s *UserService) VerifyEmail(userID uint) error {
	// Here you would implement the logic for verifying the user's email.
	return nil
}

// Implement the rest of your user service methods...
