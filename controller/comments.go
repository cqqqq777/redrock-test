package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"redrock-test/dao/mysql"
	g "redrock-test/global"
	"redrock-test/model"
	"redrock-test/services"
	"redrock-test/utils"
	"strconv"
)

// CommentBook 书评
func CommentBook(c *gin.Context) {
	comment := new(model.Comment)
	err := c.ShouldBindJSON(comment)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	comment.AuthorId = uid
	comment.Cid, err = utils.GetID()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	commentedUid, err := mysql.GetAuthorIdByBid(comment.BookId)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	comment.CommentedUid = commentedUid
	err = services.CommentBook(comment)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}

// ReplyComment 回复某一条评论
func ReplyComment(c *gin.Context) {
	comment := new(model.Comment)
	err := c.ShouldBindJSON(comment)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if comment.ParentId == 0 || comment.RootId == 0 || comment.CommentedUid == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	authorId, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	comment.AuthorId = authorId
	comment.Cid, err = utils.GetID()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	if err = services.ReplyComment(comment); err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}

// DeleteComment 删除某一条评论
func DeleteComment(c *gin.Context) {
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
	err = services.DeleteComment(cid, int64(uid))
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

// BookCommentList 获取某本书的书评
func BookCommentList(c *gin.Context) {
	pidStr := c.Param("bid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uidStr := c.Request.Header.Get("uid")
	var uid int64
	if uidStr == "" {
		uid = 0
	} else {
		uid, err = strconv.ParseInt(uidStr, 10, 64)
	}
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.BookCommentList(pid, uid, page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}

// ReplyList 获取书评的回复
func ReplyList(c *gin.Context) {
	cidStr := c.Param("cid")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uidStr := c.Request.Header.Get("uid")
	var uid int64
	uid, _ = strconv.ParseInt(uidStr, 10, 64)
	page, size := utils.GetPageInfo(c)
	data, err := services.ReplyList(cid, uid, page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}
