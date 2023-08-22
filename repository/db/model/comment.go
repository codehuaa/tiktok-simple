package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoID    uint   `gorm:"index:idx_videoid;not null" json:"video_id"`
	UserID     uint   `gorm:"index:idx_userid;not null" json:"user_id"`
	Content    string `gorm:"type:varchar(255);not null" json:"content"`
	LikeCount  uint   `gorm:"column:like_count;default:0;not null" json:"like_count,omitempty"`
	TeaseCount uint   `gorm:"column:tease_count;default:0;not null" json:"tease_count,omitempty"`
}
