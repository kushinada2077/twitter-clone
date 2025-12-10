package services

import (
	"errors"
	"twitter-clone/pkg/domain/domainerrors"
	"twitter-clone/pkg/models"
	"twitter-clone/repositories"
	"twitter-clone/utils"

	"gorm.io/gorm"
)

type AuthService interface {
	Signup(username, password string) (uint, error)
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(u repositories.UserRepository) AuthService {
	return &authService{
		userRepo: u,
	}
}

func (s *authService) Signup(username, password string) (uint, error) {
	if username == "" || password == "" {
		return 0, domainerrors.ErrEmptyUsernameOrPassword
	}

	if len(password) < 6 {
		return 0, domainerrors.ErrShortPasswordLength
	}

	exists, err := s.userRepo.Exists(username)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, domainerrors.ErrUsernameAlredayExists
	}

	hashed, err := utils.HashPassWord(password)
	if err != nil {
		return 0, err
	}

	u := &models.User{Username: username, PasswordHash: hashed}
	if u, err = s.userRepo.Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (s *authService) Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", domainerrors.ErrEmptyUsernameOrPassword
	}

	u, err := s.userRepo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", domainerrors.ErrInvalidUsernameOrPassword
		}
		return "", err
	}

	if err := utils.ComparePassword(u.PasswordHash, password); err != nil {
		return "", domainerrors.ErrInvalidUsernameOrPassword
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
