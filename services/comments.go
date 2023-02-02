package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/model"
)

func CommentBook(comment *model.Comment) (err error) {
	return mysql.CommentBook(comment)
}
