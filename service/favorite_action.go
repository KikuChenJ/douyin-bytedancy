package service

import (
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type FavoriteActionFlow struct {
	UserId     int
	Token      string
	VideoId    int
	ActionType string
}

func FavoriteAction(userId int, token string, videoId int, actionType string) error {
	return NewFavoriteActionFlow(userId, token, videoId, actionType).Do()
}

func NewFavoriteActionFlow(userId int, token string, videoId int, actionType string) *FavoriteActionFlow {
	return &FavoriteActionFlow{
		UserId:     userId,
		Token:      token,
		VideoId:    videoId,
		ActionType: actionType,
	}
}

func (f *FavoriteActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	if err := f.packData(); err != nil {
		return err
	}
	return nil
}

func (f *FavoriteActionFlow) checkParam() error {
	return nil
}

//前端传来的userId为0
func (f *FavoriteActionFlow) prepareData() error {
	favoriteDao := models.NewFavoriteDaoInstance()
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	favorite := models.Favorite{
		UserId:  userId,
		VideoId: f.VideoId,
	}

	if f.ActionType == "1" {
		err = favoriteDao.CreateFavorite(favorite)
	} else {
		err = favoriteDao.DeleteFavorite(favorite)
	}
	if err != nil {
		return err
	}
	return nil
}

func (f *FavoriteActionFlow) packData() error {

	return nil
}
