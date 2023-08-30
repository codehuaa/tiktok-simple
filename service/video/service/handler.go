package service

import (
	"context"
	"tiktok-simple/kitex/kitex_gen/user"
	video "tiktok-simple/kitex/kitex_gen/video"
	"tiktok-simple/repository/db/dao"
	"time"
)

const limit = 30

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	nextTime := time.Now().UnixMilli()
	var userID int64 = -1

	// 验证token有效性
	if req.Token != "" {
		claims, err := Jwt.ParseToken(req.Token)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "token 解析错误",
			}
			return res, nil
		}
		userID = claims.Id
	}
	// 调用数据库查询 video_list
	videos, err := dao.NewVideoDao(ctx).GetVideoList(limit, &req.LatestTime)
	if err != nil {
		res := &video.FeedResponse{
			StatusCode: -1,
			StatusMsg:  "视频获取失败：服务器内部错误",
		}
		return res, nil
	}
	videoList := make([]*video.Video, 0)
	for _, r := range videos {
		author, err := dao.NewUserDao(ctx).FindUserByUserId(r.AuthorID)
		if err != nil {
			return nil, err
		}
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}
		favorite, err := dao.NewLikeDao(ctx).GetLikeVideoRelationByUserVideoID(userID, int64(r.ID))
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}

		videoList = append(videoList, &video.Video{
			Id: int64(r.ID),
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
			PlayUrl:       r.PlayUrl,
			CoverUrl:      r.CoverUrl,
			FavoriteCount: int64(r.LikeCount),
			CommentCount:  int64(r.CommentCount),
			IsFavorite:    favorite != nil,
			Title:         r.Title,
		})
	}
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	res := &video.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
		NextTime:   nextTime,
	}
	return res, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
