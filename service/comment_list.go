package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type CommentListFlow struct {
	Token    string
	VideoId  int
	Comments []models.Comment
}

func CommentList(token string, videoId int) ([]models.Comment, error) {
	return NewCommentListFlow(token, videoId).Do()
}

func NewCommentListFlow(token string, videoId int) *CommentListFlow {
	return &CommentListFlow{
		Token:   token,
		VideoId: videoId,
	}
}

func (f *CommentListFlow) Do() ([]models.Comment, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Comments, nil
}

func (f *CommentListFlow) checkParam() error {
	if _, err := jwt.JWTAuth(f.Token); err != nil {
		return err
	}
	if f.VideoId == 0 {
		return errors.New("无视频id")
	}
	return nil
}

func (f *CommentListFlow) prepareData() error {

	return nil
}

func (f *CommentListFlow) packData() error {
	commentDao := models.NewCommentDaoInstance()
	userDao := models.NewUserDaoInstance()

	f.Comments = commentDao.MQueryCommentByVideoId(f.VideoId)
	for i := range f.Comments {
		user, _ := userDao.QueryUserById(f.Comments[i].UserId)
		f.Comments[i].User = *user
	}
	return nil
}
