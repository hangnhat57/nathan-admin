package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"personal-finance-app/pkg/domain/user"
)

type Service struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterUser(email, password string) (*user.User, error) {
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

func (s *Service) LoginUser(email, password string) (string, error) {
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

func (s *Service) VerifyEmail(userID uint) error {
	// Here you would implement the logic for verifying the user's email.
	return nil
}

// VerifyToken verifies and decodes a JWT token.
func (s *Service) VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	// Parse the token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key for verification
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Verify the token
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return token, nil
}
