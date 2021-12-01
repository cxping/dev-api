package api

import (
	"DevelopApi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiInit(route *gin.RouterGroup) {
	//模拟请求获取小车数据
	route.Handle(http.MethodPost, "/gende_server/getOdomData", get_odom_data)
	route.Handle(http.MethodGet, "/gende_server/getOdomData", get_odom_data)
}

func get_odom_data(c *gin.Context) {
	c.JSON(http.StatusOK, model.Gen_data())
}
