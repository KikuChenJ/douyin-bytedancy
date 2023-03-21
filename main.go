package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hjk-cloud/douyin/cache"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/routers"
)

func main() {
	models.Init()

	r := gin.Default()

	routers.InitRouter(r)

	////用来测试redis连接
	//cache.GetAll()

	r.Run()
}
