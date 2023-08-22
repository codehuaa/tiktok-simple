//
// Package dao
// @Description: 数据库数据库操作业务逻辑
// @Author 星梦奇缘
// @Date 2023-08-19 22:32:47
// @Update
//

package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok-simple/repository/db"
	"tiktok-simple/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao
//
//	@Description: 获取用户数据访问对象
//	@Date 2023-08-19 22:47:54
//	@param ctx 上下文
//	@return *UserDao 用户数据访问对象
func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{db.NewDBClient(ctx)}
}

// FindUserByUserName
//
//	@Description: 根据用户username获取用户数据
//	@Date 2023-08-19 22:39:54
//	@param userName 用户名
//	@return *User 用户数据
//	@return error
func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error

	return
}

// FindUserByUserId
//
//	@Description: 根据用户id获取用户数据
//	@Date 2023-08-19 10:38:54
//	@param userID 用户id
//	@return *User 用户数据
//	@return error
func (dao *UserDao) FindUserByUserId(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).
		First(&user).Error

	return
}

// CreateUser
//
//	@Description: 新增一条用户数据
//	@Date 2023-08-19 22:40:43
//	@param user 用户数据
//	@return error
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error

	return
}
