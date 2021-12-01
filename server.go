package main

import (
	"DevelopApi/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化sqlite数据库
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"code": 200, "msg": "请求成功", "data": "pong"})
	})
	//api 请求路径
	api_v1 := router.Group("/api")
	api.ApiInit(api_v1)
	router.Run(":4000")
}
