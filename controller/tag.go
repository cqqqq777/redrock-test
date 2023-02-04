package controller

import (
	"github.com/gin-gonic/gin"
	g "redrock-test/global"
	"redrock-test/services"
)

// TagList 获取标签列表
func TagList(c *gin.Context) {
	data, err := services.TagList()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}
