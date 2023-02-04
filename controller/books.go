package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"redrock-test/dao/mysql"
	g "redrock-test/global"
	"redrock-test/services"
	"redrock-test/utils"
	"strconv"
)

// TagBookList 获取某一标签下的书籍
func TagBookList(c *gin.Context) {
	tidStr := c.Param("tid")
	tid, _ := strconv.ParseInt(tidStr, 10, 64)
	if tid == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.TagBookList(tid, page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}

// BookDetail 获取某本书的详情
func BookDetail(c *gin.Context) {
	bidStr := c.Param("bid")
	bid, _ := strconv.ParseInt(bidStr, 10, 64)
	if bid == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uidStr := c.Request.Header.Get("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	data, err := services.BookDetail(bid, uid)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}

// CollectBook 收藏某本书
func CollectBook(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	bidStr := c.Param("bid")
	bid, err := strconv.ParseInt(bidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
	}
	err = services.CollectBook(bid, int64(uid))
	if err != nil {
		if errors.Is(err, mysql.ErrorRepeatOperate) {
			RespFailed(c, CodeRepeatOperate)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}

// CancelCollectBook 取消收藏
func CancelCollectBook(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	bidStr := c.Param("bid")
	bid, err := strconv.ParseInt(bidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	err = services.CancelCollectBook(bid, int64(uid))
	if err != nil {
		if errors.Is(err, mysql.ErrorNoPermission) {
			RespFailed(c, CodeNoPermission)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}

// SearchBook 搜索某本书
func SearchBook(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.SearchBooks(page, size, key)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}

// BookList 书籍列表
func BookList(c *gin.Context) {
	page, size := utils.GetPageInfo(c)
	data, err := services.BookList(page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}

// StarBook 点赞某本书
func StarBook(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	bidStr := c.Param("bid")
	bid, err := strconv.ParseInt(bidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if err = services.StarBook(bid, int64(uid)); err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}
