package service

import (
	"errors"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/util/jwt"
)

type RelationActionFlow struct {
	Token      string
	UserId     int
	ToUserId   int
	ActionType string
}

func NewRelationActionFlow(token string, userId int, toUserId int, actionType string) *RelationActionFlow {
	return &RelationActionFlow{
		Token:      token,
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: actionType,
	}
}

func RelationAction(token string, userId int, toUserId int, actionType string) error {
	return NewRelationActionFlow(token, userId, toUserId, actionType).Do()
}

func (f *RelationActionFlow) Do() error {
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

func (f *RelationActionFlow) checkParam() error {
	//fmt.Println("service---Token----", f.Token)
	//fmt.Println("service---UserId----", f.UserId)
	//fmt.Println("service---ToUserId----", f.ToUserId)
	//fmt.Println("service---ActionType----", f.ActionType)
	userId, err := jwt.JWTAuth(f.Token)
	if err != nil {
		return err
	}
	//从前端获取到的user_id一直为0，目前解决方法是根据token获取当前用户user_id
	f.UserId = userId
	if f.UserId == f.ToUserId {
		return errors.New("咋地还想关注自己啊")
	}
	return nil
}

func (f *RelationActionFlow) prepareData() error {
	relationDao := models.NewRelationDaoInstance()
	relation := models.Relation{
		UserId:   f.UserId,
		ToUserId: f.ToUserId,
	}
	if f.ActionType == "1" {
		if err := relationDao.CreateRelation(relation); err != nil {
			return err
		}
	} else if f.ActionType == "2" {
		if !relationDao.QueryRelationState(f.UserId, f.ToUserId) {
			return errors.New("未关注")
		}
		if err := relationDao.DeleteRelation(relation); err != nil {
			return err
		}
	}
	return nil
}

func (f *RelationActionFlow) packData() error {

	return nil
}
