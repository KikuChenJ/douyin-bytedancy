package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type FollowListFlow struct {
	Token  string
	UserId int
	Users  []models.User
}

func RelationFollowList(token string, userId int) ([]models.User, error) {
	return NewRelationFollowListFlow(token, userId).Do()
}

func NewRelationFollowListFlow(token string, userId int) *FollowListFlow {
	return &FollowListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *FollowListFlow) Do() ([]models.User, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f.Users, nil
}

func (f *FollowListFlow) checkParam() error {
	//fmt.Println("service---Token----", f.Token)
	//fmt.Println("service---UserId----", f.UserId)
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	if userId == 0 {
		return errors.New("这真可怕")
	}
	return nil
}

func (f *FollowListFlow) prepareData() error {

	return nil
}

func (f *FollowListFlow) packData() error {
	relationDao := models.NewRelationDaoInstance()
	userDao := models.NewUserDaoInstance()

	toUserIds := relationDao.QueryRelationByUserId(f.UserId)
	f.Users = userDao.MQueryUserById(toUserIds)
	for i := range f.Users {
		f.Users[i].IsFollow = relationDao.QueryRelationState(f.UserId, toUserIds[i])
	}
	return nil
}
