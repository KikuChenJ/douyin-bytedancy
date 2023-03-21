package models

import (
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Favorite struct {
	UserId     int       `json:"user_id"`
	VideoId    int       `json:"video_id"`
	ActionTime time.Time `json:"action_time"`
}

func (Favorite) tableName() string {
	return "favorite"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

//点赞
func (*FavoriteDao) CreateFavorite(favorite Favorite) error {
	var c chan time.Time
	c = make(chan time.Time, 10)
	c <- time.Now().UTC().Truncate(time.Second)
	for len(c) != 0 {
		//fmt.Println("len c = ", len(c))
		favorite.ActionTime = <-c
		if err := db.Table("favorite").Create(&favorite).Error; err != nil {
			return err
		}
	}
	return nil
}

//取消点赞
func (*FavoriteDao) DeleteFavorite(favorite Favorite) error {
	if err := db.Table("favorite").Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).Delete(&favorite).Error; err != nil {
		return err
	}
	return nil
}

//用户点赞数
func (*FavoriteDao) QueryFavoriteCount(videoId int) (int, error) {
	var count int64
	err := db.Table("favorite").Where("video_id = ?", videoId).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		util.Logger.Error("find relations by user_id error:" + err.Error())
		return 0, err
	}
	return int(count), nil
}

//根据userId查找已点赞视频id列表
func (*FavoriteDao) QueryFavoriteVideo(userId int) []int {
	videoIds := make([]int, 0)
	err := db.Table("favorite").Select("video_id").Order("action_time desc").Where("user_id = ?", userId).Find(&videoIds).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find favorite by user_id error:" + err.Error())
		return nil
	}
	return videoIds
}

//点赞状态 已点赞--true 未点赞--false
func (*FavoriteDao) QueryFavoriteState(userId, videoId int) bool {
	var count int64
	db.Table("favorite").Model(Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

//根据视频id查询视频的点赞数
func (*FavoriteDao) QueryVideoFavoriteCount(videoId int) int {
	var count int64
	db.Table("favorite").Model(Favorite{}).Where("video_id = ?", videoId).Count(&count)
	return int(count)
}

//根据用户id查询喜欢总数量
func (*FavoriteDao) QueryUserFavoriteCount(userId int) int {
	var count int64
	db.Table("favorite").Model(Favorite{}).Where("user_id = ?", userId).Count(&count)
	return int(count)
}
