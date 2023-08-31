package service

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"tiktok-simple/kitex/kitex_gen/user"
	video "tiktok-simple/kitex/kitex_gen/video"
	"tiktok-simple/repository/db/dao"
	"tiktok-simple/repository/db/model"
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
	// 解析token,获取用户id
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.Id

	if len(req.Title) == 0 || len(req.Title) > 32 {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "标题不能为空且不能超过32个字符",
		}
		return res, nil
	}

	createTimestamp := time.Now().UnixMilli()
	videoTitle, coverTitle := fmt.Sprintf("%d_%s_%d.mp4", userID, req.Title, createTimestamp), fmt.Sprintf("%d_%s_%d.png", userID, req.Title, createTimestamp)

	// 插入数据库
	v := &model.Video{
		Title:    req.Title,
		PlayUrl:  videoTitle,
		CoverUrl: coverTitle,
		AuthorID: uint(userID),
	}
	err = dao.NewVideoDao(ctx).CreateVideo(v)
	if err != nil {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "视频发布失败，服务器内部错误",
		}
		return res, nil
	}

	go func() {
		err := VideoPublish(req.Data, videoTitle, coverTitle)
		if err != nil {
			// 发生错误，则删除插入的记录
			e := dao.NewVideoDao(ctx).DelVideoByID(int64(v.ID), userID)
			if e != nil {
				logger.Errorf("视频记录删除失败：%s")
			}
		}
	}()
	//async.Exec(func() interface{} {
	//	return VideoPublish(ctx, req.Data, videoTitle, coverTitle, int64(v.ID), userID)
	//})
	//future.Await()

	res := &video.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "创建记录成功，等待后台上传完成",
	}
	return res, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	userID := req.UserId

	results, err := dao.NewVideoDao(ctx).GetVideosByUserID(userID)
	if err != nil {
		res := &video.PublishListResponse{
			StatusCode: -1,
			StatusMsg:  "发布列表获取失败：服务器内部错误",
		}
		return res, nil
	}
	videos := make([]*video.Video, 0)
	for _, r := range results {
		author, err := dao.NewUserDao(ctx).FindUserByUserId(r.AuthorID)
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "发布列表获取失败：服务器内部错误",
			}
			return res, nil
		}

		favorite, err := dao.NewLikeDao(ctx).GetLikeVideoRelationByUserVideoID(userID, int64(r.ID))
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "发布列表获取失败：服务器内部错误",
			}
			return res, nil
		}

		videos = append(videos, &video.Video{
			Id: int64(r.ID),
			Author: &user.User{
				Id:              int64(author.ID),
				Name:            author.UserName,
				FollowerCount:   int64(author.FollowerCount),
				FollowCount:     int64(author.FollowingCount),
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

	res := &video.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
	}
	return res, nil
}
