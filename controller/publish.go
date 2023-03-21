package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []models.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = service.Publish(token, title, data, c)

	if err == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "uploaded successfully",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}

}

//王硕-------------------------------------这个是模仿其他的List写的，通过token获取视频列表
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)

	videos, err := service.PublishList(token, userId)

	//fmt.Println(videos)
	if err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
