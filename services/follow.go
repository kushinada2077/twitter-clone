package services

import (
	"twitter-clone/pkg/domain/domainerrors"
	"twitter-clone/repositories"
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
		return domainerrors.ErrCannotFollowYourself
	}

	if _, err := s.userRepo.GetByID(followeeID); err != nil {
		return domainerrors.ErrUserNotFound
	}

	ok, err := s.followRepo.Exists(followerID, followeeID)
	if err != nil {
		return err
	}
	if ok {
		return domainerrors.ErrAlreadyFollowing
	}

	if err := s.followRepo.Create(followerID, followeeID); err != nil {
		return err
	}
	return nil
}

func (s *followService) Unfollow(followerID, followeeID uint) error {
	if followerID == followeeID {
		return domainerrors.ErrCannotUnfollowYourself
	}

	if _, err := s.userRepo.GetByID(followeeID); err != nil {
		return domainerrors.ErrUserNotFound
	}

	ok, err := s.followRepo.Exists(followerID, followeeID)
	if err != nil {
		return err
	}
	if !ok {
		return domainerrors.ErrFollowNotFound
	}

	if err := s.followRepo.Delete(followerID, followeeID); err != nil {
		return err
	}

	return nil
}
