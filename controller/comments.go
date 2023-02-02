package controller

import (
	"github.com/gin-gonic/gin"
	"redrock-test/dao/mysql"
	g "redrock-test/global"
	"redrock-test/model"
	"redrock-test/services"
	"redrock-test/utils"
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
