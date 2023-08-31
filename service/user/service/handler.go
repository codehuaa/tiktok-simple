package service

import (
	"context"
	"errors"
	user "tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/pkg/jwt"
	"tiktok-simple/repository/db/dao"
	"tiktok-simple/repository/db/model"
	"tiktok-simple/repository/redis"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// return &user.UserRegisterResponse{StatusCode: 0}, err
	// 注册
	userDao := dao.NewUserDao(ctx)
	if err = userDao.CreateUser(&model.User{
		UserName: req.Username,
		Password: req.Password,
	}); err != nil {
		// 状态码非0，表示失败
		return nil, err
	}

	// 查询 userId
	User, err := userDao.FindUserByUserName(req.Username)
	if err != nil {
		return &user.UserRegisterResponse{StatusCode: 1}, err
	}

	// 如果成功，生成token，并返回给用户
	claims := jwt.CustomClaims{
		Id: int64(User.ID),
	}
	// claims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()
	token, err := Jwt.CreateToken(claims)
	if err != nil {
		return nil, err
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
		// 如果成功，生成token，并返回给用户
		claims := jwt.CustomClaims{
			Id: int64(User.ID),
		}
		// claims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()
		var token string
		token, err = Jwt.CreateToken(claims)
		if err != nil {
			return nil, err
		}

		// 在 redis 中创建登录记录, 默认7天
		redis_client := redis.GetRedisClietn(ctx)
		err = redis_client.Set(token, 1, time.Hour*24*7).Err()
		if err != nil {
			return nil, err
		}

		return &user.UserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "success",
			UserId:     int64(User.ID),
			Token:      token,
		}, nil
	}
	return nil, errors.New("用户名密码不正确")
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	userId := req.UserId
	token := req.Token
	userDao := dao.NewUserDao(ctx)

	// 如果 id 和 token 不对应，返回
	claims, err := Jwt.ParseToken(token)
	if claims.Id != userId {
		return nil, errors.New("token和id不符！")
	}

	// 调用redis确认用户登陆状态 (即 redis 中是否存在该token)
	redis_client := redis.GetRedisClietn(ctx)
	is_exsist, err := redis_client.Exists(token).Result()
	if err != nil {
		return nil, err
	}
	if is_exsist != 1 {
		return nil, errors.New("该用户还未登录！")
	}

	// 获取用户信息
	userInfo, err := userDao.FindUserByUserId(uint(userId))
	if err != nil {
		return nil, err
	}
	return &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Id:              int64(userInfo.ID),
			Name:            userInfo.UserName,
			FollowCount:     int64(userInfo.FollowerCount),
			FollowerCount:   int64(userInfo.FollowerCount),
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			WorkCount:       int64(userInfo.WorkCount),
			FavoriteCount:   int64(userInfo.FavoritedCount),
		},
	}, nil
}
