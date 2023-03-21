package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []*models.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed @Feed
// @Description Feed
// @Accept  json
// @Produce json
// @Param   name     path    string     true        "name"
// @Success 200 {string} string	"name,helloWorld"
// @Router /api/v1/sayHello/{name} [get]
func Feed(c *gin.Context) {
	token, flag := c.GetQuery("token")
	timeStamp := c.Query("latest_time")

	var latestTime time.Time
	//fmt.Println("feed.go : Timestamp ", timeStamp)
	times, err := strconv.ParseInt(timeStamp, 10, 64)
	//fmt.Println("feed.go : times ", times, "date : Times ", time.Unix(times/1000, 0))
	if err == nil {
		latestTime = time.Unix(0, times*1e6).Local()
	}
	var videos []*models.Video

	if !flag {
		token = "-"
	}
	var nextTime int64
	videos, err, nextTime = service.VideoListWithToken(token, latestTime)

	//fmt.Println("feed.go : nextTime ", nextTime)
	//fmt.Println("controller ----------", videos)
	if err == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
