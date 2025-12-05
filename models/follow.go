package models

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	Follower  uint           `gorm:"primaryKey;autoIncrement:false" json:"follower_id"`
	Followee  uint           `gorm:"primaryKey;autoIncrement:false" json:"followee_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
