package models

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	FollowerID uint           `gorm:"primaryKey;autoIncrement:false" json:"follower_id"`
	FolloweeID uint           `gorm:"primaryKey;autoIncrement:false" json:"followee_id"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
