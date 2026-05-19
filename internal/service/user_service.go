package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/repository"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/apperror"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUsers(pageParam string, limitParam string, search string, role string) (*model.PaginationResponse, error) {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	search = strings.TrimSpace(search)
	role = strings.TrimSpace(strings.ToLower(role))

	if role != "" && role != "admin" && role != "user" {
		return nil, apperror.BadRequest("role must be admin or user")
	}
	users, total, err := s.UserRepo.FindAllPaginated(limit, offset, search, role)
	if err != nil {
		return nil, err
	}

	result := make([]model.UserResponse, 0, len(users))
	for _, user := range users {
		result = append(result, model.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &model.PaginationResponse{
		Items:      result,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

func (s *UserService) GetUserByID(idParam string) (*model.UserResponse, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return nil, apperror.BadRequest("invalid user id")
	}

	user, err := s.UserRepo.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("user not found")
		}
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserService) DeleteUser(idParam string, currentUserID uint) error {
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return apperror.BadRequest("invalid user id")
	}

	targetUserID := uint(id)

	if targetUserID == currentUserID {
		return apperror.Forbidden("you cannot delete your own account")
	}

	user, err := s.UserRepo.FindByID(targetUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("user not found")
		}
		return err
	}

	if user.Role == "admin" {
		return apperror.Forbidden("admin user cannot be deleted")
	}

	return s.UserRepo.Delete(user)
}

func (s *UserService) UpdateUser(idParam string, name string, role string) (*model.UserResponse, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return nil, apperror.BadRequest("invalid user id")
	}

	name = strings.TrimSpace(name)
	role = strings.TrimSpace(strings.ToLower(role))

	if name == "" {
		return nil, apperror.BadRequest("name is required")
	}

	if role == "" {
		role = "user"
	}

	if role != "user" && role != "admin" {
		return nil, apperror.BadRequest("role must be user or admin")
	}

	user, err := s.UserRepo.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("user not found")
		}
		return nil, err
	}

	user.Name = name
	user.Role = role

	if err := s.UserRepo.Update(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
