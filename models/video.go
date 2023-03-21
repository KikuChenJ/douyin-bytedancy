package models

import (
	"github.com/hjk-cloud/douyin/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
	"time"
)

type Video struct {
	Id            int       `gorm:"column:id;type:int" json:"id,omitempty"`
	AuthorId      int       `gorm:"column:author_id" json:"author_id,omitempty"`
	Author        User      `json:"author"`
	PlayUrl       string    `gorm:"column:play_url;type:varchar(255)" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url;type:varchar(255)" json:"cover_url"`
	FavoriteCount int       `json:"favorite_count,omitempty"`
	CommentCount  int       `json:"comment_count,omitempty"`
	Title         string    `json:"title,omitempty"`
	IsFavorite    bool      `json:"is_favorite"`
	SubmitTime    time.Time `gorm:"column:submit_time;type:datetime" json:"submit_time"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) BuildAuthor(video *Video) error {
	user, _ := NewUserDaoInstance().QueryUserById(video.AuthorId)
	video.Author = *user
	return nil
}

func (*VideoDao) MQueryVideo(videos *[]*Video, lastTime time.Time, nextTime *int64) error {
	result := db.Order("submit_time desc").Limit(30).Find(&videos, "submit_time < ?", lastTime)
	err := result.Error
	videoCnt := result.RowsAffected
	if videoCnt == 0 {
		*nextTime = lastTime.Unix() * 1000
	} else {
		*nextTime = NewVideoDaoInstance().MQueryVideoSubmitTimeById((*videos)[videoCnt-1].Id) * 1000
	}
	if err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (*VideoDao) MQueryVideoByIds(videoIds []int) []*Video {
	var videos []*Video
	err := db.Where("id in ?", videoIds).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{videoIds}, WithoutParentheses: true}}).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return nil
	}
	for i := range videos {
		NewVideoDaoInstance().BuildAuthor(videos[i])
	}
	return videos
}

func (*VideoDao) MQueryVideoSubmitTimeById(id int) int64 {
	var submitTime time.Time
	err := db.Table("Video").Select("submit_time").Where("id = ?", id).Find(&submitTime).Error
	if err != nil {
		return -1
	}
	//fmt.Println("video.go : submitTime &&  submitTime.UNix()", submitTime, submitTime.Unix())
	return submitTime.Unix()
}

//王硕-------------------通过视频id查找并返回对应的所有视频
func (*VideoDao) MQueryVideoByAuthorIds(videoIds []int) []Video {
	var videos []Video
	err := db.Where("id in ?", videoIds).Order("submit_time desc").Find(&videos).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find videos by ids error:" + err.Error())
		return nil
	}
	return videos
}

//王硕------------------------通过用户id查找该id下发布的所有视频的id
func (*VideoDao) QueryPublishVideoList(UserId int) []int {
	ids := make([]int, 0)
	err := db.Table("video").Select("id").Where("author_id = ?", UserId).Order("submit_time desc").Find(&ids).Error
	//fmt.Println(ids)
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		util.Logger.Error("find videoIds by author_id error:" + err.Error())
		return nil
	}
	return ids
}

func (*VideoDao) PublishVideo(video *Video) error {
	err := db.Select("author_id", "play_url", "cover_url", "title", "submit_time").Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
