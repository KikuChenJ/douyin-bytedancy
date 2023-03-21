package service

import (
	"github.com/hjk-cloud/douyin/models"
)

type FavoriteListFlow struct {
	Token  string
	UserId int
	Videos []*models.Video
}

func FavoriteList(token string, userId int) ([]*models.Video, error) {
	return NewFavoriteListFlow(token, userId).Do()
}

func NewFavoriteListFlow(token string, userId int) *FavoriteListFlow {
	return &FavoriteListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *FavoriteListFlow) Do() ([]*models.Video, error) {
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
func (f *FavoriteListFlow) checkParam() error {
	//fmt.Println("favoriteService---UserId----", f.UserId)

	return nil
}

func (f *FavoriteListFlow) prepareData() error {

	return nil
}

func (f *FavoriteListFlow) packData() error {
	favoriteDao := models.NewFavoriteDaoInstance()
	videoDao := models.NewVideoDaoInstance()
	relationDao := models.NewRelationDaoInstance()

	videoIds := favoriteDao.QueryFavoriteVideo(f.UserId)
	f.Videos = videoDao.MQueryVideoByIds(videoIds)

	for i := range f.Videos {
		videoDao.BuildAuthor(f.Videos[i])
		f.Videos[i].Author.IsFollow = relationDao.QueryRelationState(f.UserId, f.Videos[i].AuthorId)
		f.Videos[i].IsFavorite = favoriteDao.QueryFavoriteState(f.UserId, f.Videos[i].Id)
		f.Videos[i].FavoriteCount = favoriteDao.QueryVideoFavoriteCount(f.Videos[i].Id)
		//fmt.Println("service----Videos[i]---", f.Videos[i])
	}

	return nil
}
