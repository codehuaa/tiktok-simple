//
// Package dao
// @Description: 数据库数据库操作业务逻辑
// @Author 星梦奇缘
// @Date 2023-08-19 22:44:47
// @Update
//

package dao

import (
	"adousheng/repository/db"
	"adousheng/repository/db/model"
	"context"
	"gorm.io/gorm"
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
func (dao *videoDao) FindVideoByVideoId(id int64) (video []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("id=?", uint(id)).
		Find(&video).Error

	return
}
