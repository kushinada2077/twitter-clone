package services

import (
	"errors"
	"twitter-clone/repositories"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrFollowNotFound         = errors.New("follow not found")
	ErrAlreadyFollowing       = errors.New("already following")
	ErrCannotFollowYourself   = errors.New("cannot follow yourself")
	ErrCannotUnfollowYourself = errors.New("cannot unfollow yourself")
)

type FollowService interface {
	Follow(followerID, followeeID uint) error
	Unfollow(followerID, followeeID uint) error
}

type followService struct {
	followRepo repositories.FollowRepository
	userRepo   repositories.UserRepository
}

func NewFollowService(r repositories.FollowRepository, u repositories.UserRepository) FollowService {
	return &followService{
		followRepo: r,
		userRepo:   u,
	}
}

func (s *followService) Follow(followerID, followeeID uint) error {
	if followerID == followeeID {
		return ErrCannotFollowYourself
	}

	if _, err := s.userRepo.GetByID(followeeID); err != nil {
		return ErrUserNotFound
	}

	ok, err := s.followRepo.Exists(followerID, followeeID)
	if err != nil {
		return err
	}
	if ok {
		return ErrAlreadyFollowing
	}

	if err := s.followRepo.Create(followerID, followeeID); err != nil {
		return err
	}
	return nil
}

func (s *followService) Unfollow(followerID, followeeID uint) error {
	if followerID == followeeID {
		return ErrCannotUnfollowYourself
	}

	if _, err := s.userRepo.GetByID(followeeID); err != nil {
		return ErrUserNotFound
	}

	ok, err := s.followRepo.Exists(followerID, followeeID)
	if err != nil {
		return err
	}
	if !ok {
		return ErrFollowNotFound
	}

	if err := s.followRepo.Delete(followerID, followeeID); err != nil {
		return err
	}

	return nil
}
