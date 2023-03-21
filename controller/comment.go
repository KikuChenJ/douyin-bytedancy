package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []models.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment *models.Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	token := c.Query("token")
	videoIdString := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdString)
	actionType := c.Query("action_type") //1-发布评论，2-删除评论
	commentText := c.Query("comment_text")
	commentIdString := c.Query("comment_id")
	commentId, _ := strconv.Atoi(commentIdString)

	comment, err := service.CommentAction(userId, token, videoId, actionType, commentText, commentId)

	if err == nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0},
			Comment:  comment,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoIdString := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdString)

	comments, err := service.CommentList(token, videoId)

	if err == nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: comments,
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0, StatusMsg: err.Error()},
			CommentList: comments,
		})
	}
}
