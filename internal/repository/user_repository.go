package repository

import (
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User

	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User

	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *model.User) error {
	return r.DB.Delete(user).Error
}
func (r *UserRepository) FindAllPaginated(limit int, offset int, search string, role string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.DB.Model(&model.User{})

	if search != "" {
		searchLike := "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", searchLike, searchLike)
	}

	if role != "" {
		query = query.Where("role = ?", role)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Find(&users).Error

	return users, total, err
}
