package service

import (
	"context"
	"errors"
	favorite "tiktok-simple/kitex/kitex_gen/favorite"
	"tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/kitex/kitex_gen/video"
	"tiktok-simple/repository/db/dao"
	"tiktok-simple/repository/db/model"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: 使用redis鉴定用户登录状态
	video_id := req.VideoId
	token := req.Token
	action_type := req.ActionType
	claims, err := Jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("token错误")
	}
	user_id := claims.Id

	// 操作数据库
	likeDao := dao.NewLikeDao(ctx)
	switch action_type {
	// 1:点赞
	case 1:
		if err = likeDao.CreateVideoLike(&model.LikeVideo{uint(video_id), uint(user_id)}); err != nil {
			return
		}
	// 2:取消点赞
	case 2:
		if err = likeDao.DeleteVideoLikeByIds(video_id, user_id); err != nil {
			return
		}
	}

	return &favorite.FavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: 使用redis鉴定用户登录状态
	// token := req.Token
	user_id := req.UserId
	like_dao := dao.NewLikeDao(ctx)
	user_dao := dao.NewUserDao(ctx)

	data, err := like_dao.GetVideoLikeList(user_id)
	if err != nil {
		return
	}

	var videos_resp []*video.Video

	for _, temp_video := range data {
		// 查找user
		author, err := user_dao.FindUserByUserId(temp_video.AuthorID)
		if err != nil {
			return
		}
		videos_resp = append(videos_resp, &video.Video{
			Id: int64(temp_video.ID),
			Author: &user.User{
				Id:              int64(author.ID),
				Name:            author.UserName,
				FollowCount:     int64(author.FollowingCount),
				FollowerCount:   int64(author.FollowerCount),
				Avatar:          author.Avatar,
				BackgroundImage: author.BackgroundImage,
				Signature:       author.Signature,
				TotalFavorited:  int64(author.FavoritedCount),
				WorkCount:       int64(author.WorkCount),
				FavoriteCount:   int64(author.FavoritingCount),
			},
			PlayUrl:       temp_video.PlayUrl,
			CoverUrl:      temp_video.CoverUrl,
			FavoriteCount: int64(temp_video.LikeCount),
			CommentCount:  int64(temp_video.CommentCount),
			IsFavorite:    true,
			Title:         temp_video.Title,
		})
	}

	return &favorite.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos_resp,
	}, nil
}
