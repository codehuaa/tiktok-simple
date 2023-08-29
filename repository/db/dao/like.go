//
// Package dao
// @Description: 数据库数据库操作业务逻辑
// @Author 星梦奇缘
// @Date 2023-08-19 23:06:03
// @Update
//

package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok-simple/repository/db"
	"tiktok-simple/repository/db/model"
)

type LikeDao struct {
	*gorm.DB
}

// NewLikeDao
//
//	@Description: 获取点赞数据访问对象
//	@Date 2023-08-19 23:08:34
//	@param ctx 上下文
//	@return *LikeDao 喜欢数据访问对象
func NewLikeDao(ctx context.Context) *LikeDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &LikeDao{db.NewDBClient(ctx)}
}

// CreateVideoLike
//
//	@Description: 创建新点赞记录
//	@Date 2023-08-20 16:11:34
//	@param likeVideo 喜欢对象
//	@return err 错误信息
func (dao *LikeDao) CreateVideoLike(likeVideo *model.LikeVideo) error {
	videoId := likeVideo.VideoID
	err := dao.DB.Model(&model.LikeVideo{}).Create(likeVideo).Error
	if err != nil {
		return err
	}

	// 更新视频的FavoriteCount
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(videoId)).Update("like_count",
		gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// 更新点赞者的FavoritingCount
	user_id := likeVideo.UserID
	err = dao.DB.Model(&model.User{}).Where("id=?", uint(user_id)).Update("favoriting_count",
		gorm.Expr("favoriting_count + ?", 1)).Error
	if err != nil {
		return err
	}

	// 更新创作者的FavoritedCount
	var video *model.Video
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(videoId)).First(&video).Error
	updater_id := video.AuthorID
	if err != nil {
		return err
	}
	err = dao.DB.Model(&model.User{}).Where("id=?", updater_id).Update("favoriting_count",
		gorm.Expr("favoriting_count - ?", 1)).Error
	return err
}

// CreateCommentLike
//
//	@Description: 创建新评论点赞记录
//	@Date 2023-08-20 16:21:34
//	@param likeComment 喜欢对象
//	@return err 错误信息
func (dao *LikeDao) CreateCommentLike(likeComment *model.LikeComment) error {
	commentId := likeComment.CommentID
	err := dao.DB.Model(&model.LikeComment{}).Create(likeComment).Error
	if err != nil {
		return err
	}

	// 更新视频的FavoriteCount
	err = dao.DB.Model(&model.Comment{}).Where("id=?", uint(commentId)).Update("like_count",
		gorm.Expr("like_count + ?", 1)).Error
	return err
}

// DeleteVideoLikeByIds
//
//	@Description: 根据Id撤销视频喜欢
//	@Date 2023-08-20 16:25:54
//	@param videoLikeId
//	@return err 错误信息
func (dao *LikeDao) DeleteVideoLikeByIds(videoId int64, userId int64) error {
	var videoLike *model.LikeVideo
	err := dao.DB.Model(&model.LikeVideo{}).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&videoLike).Error
	if err != nil {
		return err
	}

	// 更新video的LikeCount属性
	video_id := videoLike.VideoID
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(video_id)).Update("like_count",
		gorm.Expr("like_count - ?", 1)).Error
	if err != nil {
		return err
	}

	// 更新点赞者的FavoritingCount
	user_id := videoLike.UserID
	err = dao.DB.Model(&model.User{}).Where("id=?", uint(user_id)).Update("favoriting_count",
		gorm.Expr("favoriting_count - ?", 1)).Error
	if err != nil {
		return err
	}

	// 更新创作者的FavoritedCount
	var video *model.Video
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(video_id)).First(&video).Error
	updater_id := video.AuthorID
	if err != nil {
		return err
	}
	err = dao.DB.Model(&model.User{}).Where("id=?", updater_id).Update("favorited_count",
		gorm.Expr("favorited_count - ?", 1)).Error
	return err

}

// DeleteCommentLikeById
//
//	@Description: 根据Id撤销评论喜欢
//	@Date 2023-08-20 19:07:34
//	@param commentLikeId
//	@return err 错误信息
func (dao *LikeDao) DeleteCommentLikeById(commentId int64, userId int64) error {
	var commentLike *model.LikeComment
	err := dao.DB.Model(&model.LikeComment{}).Where("comment_id=? and user_id=?", uint(commentId), uint(userId)).Delete(&commentLike).Error
	if err != nil {
		return err
	}

	// 更新comment的LikeCount属性
	id := commentLike.CommentID
	err = dao.DB.Model(&model.Comment{}).Where("id=?", uint(id)).Update("like_count", gorm.Expr("like_count - ?", 1)).Error
	return err
}
