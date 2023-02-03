package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/model"
)

func CommentBook(comment *model.Comment) (err error) {
	return mysql.CommentBook(comment)
}

func ReplyComment(comment *model.Comment) error {
	return mysql.ReplyComment(comment)
}

func DeleteComment(cid, uid int64) error {
	authorId, err := mysql.GetAuthorIdByCid(cid)
	if err != nil {
		return err
	}
	if authorId != uid {
		return mysql.ErrorNoPermission
	}
	return mysql.DeleteComment(cid)
}
