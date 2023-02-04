package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/model"
)

func TagBookList(tid, page, size int64) ([]*model.Book, error) {
	return mysql.TagBookList(tid, page, size)
}

func BookDetail(bid, uid int64) (data *model.ApiBook, err error) {
	book, err := mysql.BookDetail(bid)
	data = new(model.ApiBook)
	data.Book = book
	if err != nil {
		return nil, err
	}
	started, err := mysql.GetUserStarBook(bid, uid)
	if err != nil {
		return nil, err
	}
	if started == 1 {
		data.Started = true
	} else {
		data.Started = false
	}
	data.CommentNum, err = mysql.GetBookCommentsTotalNum(bid)
	if err != nil {
		return nil, err
	}
	data.Tags, err = mysql.GetBookTags(bid)
	if err != nil {
		return nil, err
	}
	return
}

func CollectBook(bid, uid int64) (err error) {
	started, err := mysql.GetUserStarBook(bid, uid)
	if err != nil {
		return err
	}
	if started != 0 {
		return mysql.ErrorRepeatOperate
	}
	err = mysql.CollectBook(bid, uid)
	return
}

func CancelCollectBook(bid, uid int64) error {
	started, err := mysql.GetUserStarBook(bid, uid)
	if err != nil {
		return err
	}
	if started == 0 {
		return mysql.ErrorNoPermission
	}
	return mysql.CancelCollectBook(bid, uid)
}

func SearchBooks(page, size int64, key string) ([]*model.Book, error) {
	return mysql.SearchBooks(page, size, key)
}

func BookList(page, size int64) ([]*model.Book, error) {
	return mysql.BookList(page, size)
}
