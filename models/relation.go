package models

import (
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Relation struct {
	UserId   int `json:"user_id"`
	ToUserId int `json:"to_user_id"`
}

func (Relation) TableName() string {
	return "relation"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

//关注
func (*RelationDao) CreateRelation(relation Relation) error {
	if err := db.Create(&relation).Error; err != nil {
		return err
	}
	return nil
}

//取关
func (*RelationDao) DeleteRelation(relation Relation) error {
	if err := db.Where("user_id = ? and to_user_id = ?", relation.UserId, relation.ToUserId).Delete(&relation).Error; err != nil {
		return err
	}
	return nil
}

//关注列表
func (*RelationDao) QueryRelationByUserId(userId int) []int {
	ids := make([]int, 0)
	err := db.Table("relation").Select("to_user_id").Where("user_id = ?", userId).Find(&ids).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return nil
	}
	return ids
}

//关注人数
func (*RelationDao) QueryRelationCountByUserId(userId int) (int, error) {
	var count int64
	err := db.Table("relation").Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

//粉丝列表
func (*RelationDao) QueryRelationByToUserId(userId int) []int {
	ids := make([]int, 0)
	err := db.Table("relation").Select("user_id").Where("to_user_id = ?", userId).Find(&ids).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return nil
	}
	return ids
}

//粉丝人数
func (*RelationDao) QueryRelationCountByToUserId(userId int) (int, error) {
	var count int64
	err := db.Table("relation").Where("to_user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

//关注状态 已关注--true 未关注--false
func (*RelationDao) QueryRelationState(userId int, toUserId int) bool {
	var count int64
	db.Model(Relation{}).Where("user_id = ? and to_user_id = ?", userId, toUserId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
