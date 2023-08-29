package service

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"gorm.io/gorm"
	comment "tiktok-simple/kitex/kitex_gen/comment"
	kitex "tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/repository/db/dao"
	"tiktok-simple/repository/db/model"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	claims, err := Jwt.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("token解析错误：%v", err.Error())
		res := &comment.CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.Id
	actionType := req.ActionType
	videoDao := dao.NewVideoDao(ctx)
	v, _ := videoDao.FindVideoByVideoId(req.VideoId)
	if v == nil {
		logger.Errorf("该视频ID不存在：%d", req.VideoId)
		res := &comment.CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "该视频ID不存在",
		}
		return res, nil
	}
	if actionType == 1 {
		cmt := &model.Comment{
			VideoID: uint(req.VideoId),
			UserID:  uint(userID),
			Content: req.CommentText,
		}
		commentDao := dao.NewCommentDao(ctx)
		err := commentDao.CreateComment(cmt)
		if err != nil {
			logger.Errorf("新增评论失败：%v", err.Error())
			res := &comment.CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论发布失败：服务器内部错误",
			}
			return res, nil
		}
	} else if actionType == 2 {
		// 判断该评论是否发布自该用户，或该评论在该用户所发布的视频下
		commentDao := dao.NewCommentDao(ctx)
		cmt, err := commentDao.FindCommentByCommentID(req.CommentId)
		if err != nil {
			logger.Errorf("评论删除失败：%v", err.Error())
			res := &comment.CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论删除失败：服务器内部错误",
			}
			return res, nil
		}
		if cmt == nil {
			// 评论不存在，无法删除
			logger.Errorf("评论删除失败，该评论ID不存在：%v", req.CommentId)
			res := &comment.CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论删除失败：该评论不存在",
			}
			return res, nil
		} else {
			// 查找该视频的作者ID
			v, err := videoDao.FindVideoByVideoId(int64(cmt.VideoID))
			if err != nil {
				logger.Errorf("评论删除失败：%v", err.Error())
				res := &comment.CommentActionResponse{
					StatusCode: -1,
					StatusMsg:  "评论删除失败：服务器内部错误",
				}
				return res, nil
			}
			// 若删除评论的用户不是发布评论的用户或该用户不是视频创作者
			if userID != int64(cmt.UserID) || userID != int64(v.AuthorID) {
				logger.Errorf("评论删除失败，没有权限：%v", cmt.UserID)
				res := &comment.CommentActionResponse{
					StatusCode: -1,
					StatusMsg:  "评论删除失败：没有权限",
				}
				return res, nil
			}
		}
		err = commentDao.DeleteCommentByID(req.CommentId)
		if err != nil {
			logger.Errorf("评论删除失败：%v", err.Error())
			res := &comment.CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "评论删除失败：服务器内部错误",
			}
			return res, nil
		}
	} else {
		res := &comment.CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "action_type 非法",
		}
		return res, nil
	}
	res := &comment.CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return res, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	var userID int64 = -1
	// 验证token有效性
	if req.Token != "" {
		claims, err := Jwt.ParseToken(req.Token)
		if err != nil {
			logger.Errorf("token解析错误:%v", err)
			res := &comment.CommentListResponse{
				StatusCode: -1,
				StatusMsg:  "token 解析错误",
			}
			return res, nil
		}
		userID = claims.Id
	}

	// 从数据库获取评论列表
	commentDao := dao.NewCommentDao(ctx)
	results, err := commentDao.FindCommentListByVideoID(req.VideoId)
	if err != nil {
		res := &comment.CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "评论列表获取失败：服务器内部错误",
		}
		return res, nil
	}
	comments := make([]*comment.Comment, 0)
	for _, r := range results {
		userDao := dao.NewUserDao(ctx)
		u, err := userDao.FindUserByUserId(uint(userID))
		if err != nil && err != gorm.ErrRecordNotFound {
			res := &comment.CommentListResponse{
				StatusCode: -1,
				StatusMsg:  "评论列表获取失败：服务器内部错误",
			}
			return res, nil
		}
		usr := &kitex.User{
			Id:              userID,
			Name:            u.UserName,
			FollowCount:     int64(u.FollowingCount),
			FollowerCount:   int64(u.FollowerCount),
			IsFollow:        err != gorm.ErrRecordNotFound,
			Signature:       u.Signature,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			TotalFavorited:  int64(u.FavoritedCount),
			WorkCount:       int64(u.WorkCount),
			FavoriteCount:   int64(u.FavoritingCount),
		}
		comments = append(comments, &comment.Comment{
			Id:         int64(r.ID),
			User:       usr,
			Content:    r.Content,
			CreateDate: r.CreatedAt.Format("2006-01-02"),
			LikeCount:  int64(r.LikeCount),
			TeaseCount: int64(r.TeaseCount),
		})
	}
	return
}
