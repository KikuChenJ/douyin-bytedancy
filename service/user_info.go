package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
)

type UserInfoFlow struct {
	UserId         int
	Token          string
	User           *models.User
	FollowCount    int
	FollowerCount  int
	TotalFavorited int
	FavoriteCount  int
}

func UserInfo(token string, userId int) (*models.User, error) {
	return NewUserInfoFlow(token, userId).Do()
}

func NewUserInfoFlow(token string, userId int) *UserInfoFlow {
	return &UserInfoFlow{Token: token, UserId: userId}
}

func (f *UserInfoFlow) Do() (*models.User, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.User, nil
}

//此处不能验证token
//对于未登录的用户，想要查看视频作者信息时，不需要token即可查看
func (f *UserInfoFlow) checkParam() error {
	if f.UserId == 0 {
		return errors.New("id为空")
	}
	return nil
}

func (f *UserInfoFlow) prepareData() error {
	relationDao := models.NewRelationDaoInstance()
	favoriteDao := models.NewFavoriteDaoInstance()
	videoDao := models.NewVideoDaoInstance()
	var totalFavorited = 0

	//关注数
	followCount, err := relationDao.QueryRelationCountByUserId(f.UserId)
	if err != nil {
		return err
	}
	f.FollowCount = followCount
	//粉丝数
	followerCount, err := relationDao.QueryRelationCountByToUserId(f.UserId)
	if err != nil {
		return err
	}
	f.FollowerCount = followerCount
	//获赞数
	videoIds := videoDao.QueryPublishVideoList(f.UserId)
	for i := range videoIds {
		totalFavorited += favoriteDao.QueryVideoFavoriteCount(videoIds[i])
	}
	f.TotalFavorited = totalFavorited
	//喜欢数
	f.FavoriteCount = favoriteDao.QueryUserFavoriteCount(f.UserId)
	return nil
}

func (f *UserInfoFlow) packData() error {
	userDao := models.NewUserDaoInstance()

	user, err := userDao.QueryUserById(f.UserId)
	if err != nil {
		return err
	}
	f.User = user
	f.User.FollowCount = f.FollowCount
	f.User.FollowerCount = f.FollowerCount
	f.User.TotalFavorited = f.TotalFavorited
	f.User.FavoriteCount = f.FavoriteCount
	return nil
}
