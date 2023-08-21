//
// Package dao
// @Description: 数据库数据库操作业务逻辑
// @Author 星梦奇缘
// @Date 2023-08-20 11:45:03
// @Update
//

package dao

import (
	"adousheng/repository/db"
	"adousheng/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

// NewCommentDao
//
//	@Description: 获取用户数据访问对象
//	@Date 2023-08-20 11:47:03
//	@param ctx 上下文
//	@return *UserDao 评论数据访问对象
func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{db.NewDBClient(ctx)}
}

// FindCommentListByVideoID
//
//	@Description: 根据视频id获取指定视频的全部评论内容
//	@Date 2023-08-20 11:51:33
//	@param videoID 视频id
//	@return []*Comment 评论内容
//	@return error
func (dao *CommentDao) FindCommentListByVideoID(videoID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := dao.DB.Model(&model.Comment{}).Where("video_id=?", uint(videoID)).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindCommentByCommentID
//
//	@Description: 根据评论id获取评论内容
//	@Date 2023-08-20 11:57:27
//	@param commentID 评论id
//	@return *Comment 评论
//	@return error
func (dao *CommentDao) FindCommentByCommentID(commentID int64) (comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", uint(commentID)).First(&comment).Error
	return
}

// CreateComment
//
//	@Description: 获取新评论
//	@Date 2023-08-20 11:57:27
//	@param comment 评论对象
//	@return *Comment 评论
//	@return error
func (dao *CommentDao) CreateComment(comment *model.Comment) error {
	err := dao.DB.Model(&model.Comment{}).Create(comment).Error
	if err != nil {
		return err
	}

	id := comment.VideoID
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(id)).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	return err
}

// DeleteCommentByID
//
//	@Description: 根据ID删除评论
//	@Date 2023-08-20 13:28:20
//	@param commentID 评论id
//	@return error
func (dao *CommentDao) DeleteCommentByID(commentID int64) error {
	var comment *model.Comment
	err := dao.DB.Model(&model.Comment{}).Where("id=?", uint(commentID)).Delete(&comment).Error
	if err != nil {
		return err
	}

	// 更新video的CommentCount属性
	id := comment.VideoID
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(id)).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	return err
}
