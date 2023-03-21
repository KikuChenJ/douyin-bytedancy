package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
)

type FavoriteVideoListResponse struct {
	Response
	VideoList []*models.Video `json:"video_list"`
}

func FavoriteAction(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	token := c.Query("token")
	videoIdString := c.Query("video_id")
	videoId, _ := strconv.Atoi(videoIdString)
	actionType := c.Query("action_type") //1-点赞，2-取消点赞

	if err := service.FavoriteAction(userId, token, videoId, actionType); err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	}

}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)

	videos, _ := service.FavoriteList(token, userId)

	c.JSON(http.StatusOK, FavoriteVideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
