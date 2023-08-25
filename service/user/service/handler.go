package service

import (
	"context"
	user "tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/repository/db/dao"
	"tiktok-simple/repository/db/model"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// return &user.UserRegisterResponse{StatusCode: 0}, err
	// 注册
	userDao := dao.NewUserDao(ctx)
	err = userDao.CreateUser(&model.User{
		UserName: req.Username,
		Password: req.Password,
	})
	if err != nil {
		// 状态码非0，表示失败
		return nil, err
	}

	// TODO 如果成功，生成token，并返回给用户
	token := "this is tmp token"

	// 查询 userId
	User, err := userDao.FindUserByUserName(req.Username)
	if err != nil {
		return &user.UserRegisterResponse{StatusCode: 1}, err
	}

	resp = &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(User.ID),
		Token:      token,
	}
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	userDao := dao.NewUserDao(ctx)
	User, err := userDao.FindUserByUserName(req.Username)
	if err != nil {
		return nil, err
	}
	if User.Password == req.Password {
		// TODO 如果成功，生成token，并返回给用户
		token := "this is tmp token"
		return &user.UserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			UserId:     int64(User.ID),
			Token:      token,
		}, nil
	}
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
