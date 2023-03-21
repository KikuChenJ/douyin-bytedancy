package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/service"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []models.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)
	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.Atoi(toUserIdString)
	actionType := c.Query("action_type") //1-关注，2-取消关注

	err := service.RelationAction(token, userId, toUserId, actionType)

	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	}
}

//关注列表
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)

	users, err := service.RelationFollowList(token, userId)

	if err == nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

//粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdString)

	users, err := service.RelationFollowerList(token, userId)

	if err == nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
