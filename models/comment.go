package models

import (
	"gorm.io/gorm"
	"sync"
)

type Comment struct {
	Id         int    `json:"id,omitempty"`
	UserId     int    `json:"user_id"`
	User       User   `json:"user"`
	VideoId    int    `json:"video_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

//根据视频id查找相应评论
func (*CommentDao) MQueryCommentByVideoId(videoId int) []Comment {
	var comments []Comment
	err := db.Order("create_date desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return nil
	}
	return comments
}

//发布评论
func (*CommentDao) CreateComment(comment *Comment) error {
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

//删除评论
func (*CommentDao) DeleteComment(comment *Comment) error {
	if err := db.Where("id = ?", comment.Id).Delete(comment).Error; err != nil {
		return err
	}
	return nil
}

//根据videoId查询评论数
func (*CommentDao) QueryCommentCount(videoId int) int {
	var count int64
	db.Model(Comment{}).Where("video_id = ?", videoId).Count(&count)
	return int(count)
}
