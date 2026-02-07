package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/rest-api-notes/internal/config"
	models "example.com/rest-api-notes/internal/domain"
	"example.com/rest-api-notes/internal/repository"
	"github.com/golang-jwt/jwt"
	"github.com/keybase/go-crypto/bcrypt"
	"go.uber.org/zap"
)

type UserService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
	logger    *zap.Logger
}

func NewUserService(
	userRepo *repository.UserRepository,
	authconfig config.AuthConfig,
	logger *zap.Logger,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwtSecret: authconfig.JWTSecret,
		logger: logger,
	}
}

func (s *UserService) SignUp(ctx context.Context, req models.SignUpRequest) (*models.AuthResponse, error) {
	s.logger.Info("Attempting to sign up user", zap.String("email", req.Email))

	//check if email already exists
	exists, err := s.userRepo.EmailExists(ctx, req.Email)
	if err != nil {
		s.logger.Error("Failed to check email existence", zap.Error(err))
		return nil, fmt.Errorf("failed to check email: %w", err)
	}

	if exists {
		s.logger.Warn("Email already registered", zap.String("email", req.Email))
		return nil, errors.New("email already registered")
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	//create user
	user, err := s.userRepo.SignUp(ctx, req.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	//generate token
	token, err := s.generateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
	
}

func (s *UserService) generateToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}