package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type FollowerListFlow struct {
	Token  string
	UserId int
	Users  []models.User
}

func RelationFollowerList(token string, userId int) ([]models.User, error) {
	return NewRelationFollowerListFlow(token, userId).Do()
}

func NewRelationFollowerListFlow(token string, userId int) *FollowerListFlow {
	return &FollowerListFlow{
		Token:  token,
		UserId: userId,
	}
}

func (f *FollowerListFlow) Do() ([]models.User, error) {
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

func (f *FollowerListFlow) checkParam() error {
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

func (f *FollowerListFlow) prepareData() error {

	return nil
}

func (f *FollowerListFlow) packData() error {
	relationDao := models.NewRelationDaoInstance()
	userDao := models.NewUserDaoInstance()

	userIds := relationDao.QueryRelationByToUserId(f.UserId)
	f.Users = userDao.MQueryUserById(userIds)
	for i := range f.Users {
		f.Users[i].IsFollow = relationDao.QueryRelationState(f.UserId, userIds[i])
	}
	return nil
}
