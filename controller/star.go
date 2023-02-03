package controller

import (
	"github.com/gin-gonic/gin"
	g "redrock-test/global"
	"redrock-test/services"
	"redrock-test/utils"
	"strconv"
)

func StarComment(c *gin.Context) {
	cidStr := c.Param("cid")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	if err = services.StarComment(cid, int64(uid)); err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}
