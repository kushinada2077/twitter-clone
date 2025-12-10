package repositories

import (
	"twitter-clone/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Create(u *models.User) (*models.User, error)
	Exists(username string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Take(&user, "username = ?", username).Error
	return &user, err
}

func (r *userRepository) Create(u *models.User) (*models.User, error) {
	if err := r.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) Exists(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, err
}
