package service

import (
	"errors"
	"strings"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/config"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/repository"
	jwthelper "github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepo *repository.UserRepository
	Config   *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		Config:   cfg,
	}
}

func (s *AuthService) Register(req model.RegisterRequest) (*model.UserResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	existingUser, err := s.UserRepo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := s.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *AuthService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwthelper.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
		s.Config.JWTSecret,
		s.Config.JWTExpiredHour,
	)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
		User: model.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
