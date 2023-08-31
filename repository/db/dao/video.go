//
// Package dao
// @Description: 数据库数据库操作业务逻辑
// @Author 星梦奇缘
// @Date 2023-08-19 22:44:47
// @Update
//

package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok-simple/repository/db"
	"tiktok-simple/repository/db/model"
	"time"
)

type videoDao struct {
	*gorm.DB
}

// NewVideoDao
//
//	@Description: 获取视频数据访问对象
//	@Date 2023-08-19 22:50:54
//	@param ctx 上下文
//	@return *UserDao 视频数据访问对象
func NewVideoDao(ctx context.Context) *videoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &videoDao{db.NewDBClient(ctx)}
}

// FindVideoListByVideoName
// @Description: 根据视频的title模糊查询获取视频，并按照发布时间排序
// @Date 2023-08-19 22:58:54
// @param title 视频名
// @return *video 视频数据
// @return error
func (dao *videoDao) FindVideoListByVideoName(title string) (video []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("title LIKE ?", title).
		Find(&video).Error

	return
}

// FindVideoByVideoId
// @Description: 根据视频的id模糊查询获取视频
// @Date 2023-08-19 23:00:33
// @param id 视频id
// @return *video 视频数据
// @return error
func (dao *videoDao) FindVideoByVideoId(id int64) (video *model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(id)).
		Find(&video).Error

	return
}

func (dao *videoDao) GetVideoList(limit int, latestTime *int64) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)

	if latestTime == nil || *latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = &curTime
	}
	if err := dao.Limit(limit).Order("created_at desc").Find(&videos, "created_at < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (dao *videoDao) CreateVideo(video *model.Video) error {
	err := dao.Transaction(func(tx *gorm.DB) error {
		// 1. 在 video 表中创建视频记录
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		// 2. 同步 user 表中的作品数量
		res := tx.Model(&model.User{}).Where("id = ?", video.AuthorID).Update("work_count", gorm.Expr("work_count + ?", 1))
		if res.Error != nil {
			return err
		}
		if res.RowsAffected != 1 {
			return gorm.ErrInvalidField
		}
		return nil
	})

	return err
}

func (dao *videoDao) DelVideoByID(videoID int64, authorID int64) error {
	err := dao.Transaction(func(tx *gorm.DB) error {
		// 1. 根据主键 video_id 删除 video
		err := tx.Unscoped().Delete(&model.Video{}, videoID).Error
		if err != nil {
			return err
		}
		// 2. 同步 user 表中的作品数量
		res := tx.Model(&model.User{}).Where("id = ?", authorID).Update("work_count", gorm.Expr("work_count - ?", 1))
		if res.Error != nil {
			return err
		}
		if res.RowsAffected != 1 {
			return gorm.ErrInvalidField
		}
		return nil
	})
	return err
}

func (dao *videoDao) GetVideosByUserID(authorId int64) ([]*model.Video, error) {
	var pubList []*model.Video
	err := dao.Model(&model.Video{}).Where(&model.Video{AuthorID: uint(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
