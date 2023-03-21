package service

import (
	"github.com/hjk-cloud/douyin/models"
)

type PublishListFlow struct {
	Token  string
	UserId int
	Videos []models.Video
}

func PublishList(token string, userId int) ([]models.Video, error) {
	return NewPublishListWithTokenFlow(token, userId).Do()
}

func NewPublishListWithTokenFlow(token string, userId int) *PublishListFlow {
	return &PublishListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *PublishListFlow) Do() ([]models.Video, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Videos, nil
}

//此处不能验证token
//对于未登录的用户，想要查看视频作者信息时，不需要token即可查看
func (f *PublishListFlow) checkParam() error {

	return nil
}

func (f *PublishListFlow) prepareData() error {

	//fmt.Println("service-token--userId----", userId)
	//fmt.Println("service---AuthorId----", f.UserId)
	return nil
}

func (f *PublishListFlow) packData() error {
	videoDao := models.NewVideoDaoInstance()
	favoriteDao := models.NewFavoriteDaoInstance()

	videos := videoDao.QueryPublishVideoList(f.UserId)
	f.Videos = videoDao.MQueryVideoByAuthorIds(videos)
	for i := range f.Videos {
		f.Videos[i].IsFavorite = favoriteDao.QueryFavoriteState(f.UserId, f.Videos[i].Id)
		f.Videos[i].FavoriteCount = favoriteDao.QueryVideoFavoriteCount(f.Videos[i].Id)
	}
	//fmt.Println("video-----------", videos)
	return nil
}
