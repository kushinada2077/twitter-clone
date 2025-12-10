package repositories

import (
	"twitter-clone/pkg/models"

	"gorm.io/gorm"
)

type FollowRepository interface {
	Create(followerID, followeeID uint) error
	Delete(followerID, followeeID uint) error
	Exists(followerID, followeeID uint) (bool, error)
}

type followRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{
		db: db,
	}
}

func (r *followRepository) Create(followerID, followeeID uint) error {
	f := models.Follow{FollowerID: followerID, FolloweeID: followeeID}
	return r.db.Create(&f).Error
}

func (r *followRepository) Delete(followerID, followeeID uint) error {
	return r.db.Where("follower_id = ? AND followee_id = ?", followerID, followeeID).Delete(&models.Follow{}).Error
}

func (r *followRepository) Exists(followerID, followeeID uint) (bool, error) {
	var exists bool
	query := `
			SELECT EXISTS (
				SELECT 1
				FROM follows
				WHERE follower_id = ?
					AND followee_id = ?
					AND deleted_at IS NULL
	)`

	err := r.db.Raw(query, followerID, followeeID).Scan(&exists).Error
	return exists, err
}
