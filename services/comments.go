package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/dao/redisdao"
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

func BookCommentList(bid, uid, page, size int64) (data *model.ApiCommentList, err error) {
	comments, err := mysql.BookCommentList(bid, page, size)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiCommentList)
	data.Comments = make([]*model.ApiComment, 0)
	for _, comment := range comments {
		comment.Author, err = mysql.FindUsernameByUid(comment.Comment.AuthorId)
		if err != nil {
			continue
		}
		comment.ReplyNum, err = mysql.GetReplyNum(comment.Comment.Cid)
		if err != nil {
			continue
		}
		comment.Stars, err = redisdao.GetCommentsStars(int64(comment.Comment.Cid))
		if err != nil {
			comment.Stars = 0
		}
		id, err := mysql.GetIdByUid(uid)
		if err != nil {
			comment.Started = false
		}
		status, err := redisdao.GetUserStarCommentStatus(int64(comment.Comment.Cid), id)
		if err != nil {
			comment.Started = false
		}
		switch status {
		case 0:
			comment.Started = false
		case 1:
			comment.Started = true
		}
		data.Comments = append(data.Comments, comment)
	}
	data.TotalNum, err = mysql.GetBookCommentsTotalNum(bid)
	if err != nil {
		return nil, err
	}
	return
}

func ReplyList(cid, uid, page, size int64) (data []*model.ApiReply, err error) {
	replies, err := mysql.ReplyList(cid, page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*model.ApiReply, 0)
	for _, reply := range replies {
		reply.Author, err = mysql.FindUsernameByUid(reply.Comment.AuthorId)
		if err != nil {
			continue
		}
		reply.CommentedPeople, err = mysql.FindUsernameByUid(int(reply.Comment.CommentedUid))
		if err != nil {
			continue
		}
		reply.Stars, _ = redisdao.GetCommentsStars(int64(reply.Comment.Cid))
		id, err := mysql.GetIdByUid(uid)
		if err != nil {
			reply.Started = false
		}
		status, err := redisdao.GetUserStarCommentStatus(int64(reply.Comment.Cid), id)
		if err != nil {
			reply.Started = false
		}
		switch status {
		case 0:
			reply.Started = false
		case 1:
			reply.Started = true
		}
		data = append(data, reply)
	}
	return data, nil
}
