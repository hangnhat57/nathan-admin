package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"personal-finance-app/pkg/cfg"
	"personal-finance-app/pkg/domain/user"
	"personal-finance-app/pkg/infrastructure/email"
	"time"
)

type Service struct {
	repo   user.Repository
	mailer *email.Emailer
}

func NewUserService(repo user.Repository, mailer *email.Emailer) *Service {
	return &Service{repo: repo, mailer: mailer}
}

func (s *Service) RegisterUser(request user.SignUpRequest) (*user.User, error) {
	// Here you would implement the logic for user registration.
	// This could include validating the input, checking if a user with the given email already exists, etc.
	// Then, hash the password before storing it in the database.

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	// Generate a verification token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": newUser.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Set the token expiration time
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString(cfg.JwtSecret)
	if err != nil {
		return nil, err
	}

	// Send the verification email with the token
	err = s.mailer.SendEmail(newUser.Email, "Verify your email",
		fmt.Sprintf("Click on this \n %s", getVerificationLink(tokenString)))
	if err != nil {
		return nil, err
	}

	return newUser, nil
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

func getVerificationLink(token string) string {
	return fmt.Sprintf("%s%s", cfg.EmailVerificationLinkPrefix, token) // Replace with your verification endpoint URL
}
