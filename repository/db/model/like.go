package model

type LikeVideo struct {
	VideoID uint `gorm:"index:idx_videoid;not null" json:"video_id"`
	UserID  uint `gorm:"index:idx_userid;not null" json:"user_id"`
}

type LikeComment struct {
	CommentID uint `gorm:"index:idx_commentid;not null" json:"comment_id"`
	UserID    uint `gorm:"index:idx_userid;not null" json:"user_id"`
}
