package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

const (
	isDel = "JscSqqCsPsQcSYYsaPPQ"
)

func CommentBook(comment *model.Comment) (err error) {
	_, err = g.Mdb.Exec("insert into comments (cid, author_id, book_id, commented_uid, content) values (?,?,?,?,?)", comment.Cid, comment.AuthorId, comment.BookId, comment.CommentedUid, comment.Content)
	return
}

func ReplyComment(comment *model.Comment) error {
	_, err := g.Mdb.Exec("insert into comments(cid,author_id,book_id,parent_id,root_id,commented_uid,content) values(?,?,?,?,?,?,?)", comment.Cid, comment.AuthorId, comment.BookId, comment.ParentId, comment.RootId, comment.CommentedUid, comment.Content)
	return err
}

func GetAuthorIdByCid(cid int64) (authorId int64, err error) {
	err = g.Mdb.QueryRow("select author_id from comments where cid = ?", cid).Scan(&authorId)
	return
}

func DeleteComment(cid int64) (err error) {
	_, err = g.Mdb.Exec("update comments set content = ? where cid = ? or root_id = ?", isDel, cid, cid)
	return
}

func BookCommentList(bid, page, size int64) (comments []*model.ApiComment, err error) {
	comments = make([]*model.ApiComment, 0)
	rows, err := g.Mdb.Query("select cid,author_id,book_id,commented_uid,content,create_time from comments where book_id = ? and root_id = 1 and content <>? ORDER BY stars desc limit ?,?", bid, isDel, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		comment := new(model.Comment)
		err = rows.Scan(&comment.Cid, &comment.AuthorId, &comment.BookId, &comment.CommentedUid, &comment.Content, &comment.CreateTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &model.ApiComment{
			Comment: comment,
		})
	}
	return
}

func GetReplyNum(cid int) (replyNum int64, err error) {
	err = g.Mdb.QueryRow("select count(cid) from comments where  root_id = ? and content <> ?", cid, isDel).Scan(&replyNum)
	return
}

func GetBookCommentsTotalNum(bid int64) (totalNum int64, err error) {
	err = g.Mdb.QueryRow("select count(cid) from comments where book_id =? and root_id = 1 and content <> ?", bid, isDel).Scan(&totalNum)
	return
}

func ReplyList(cid, page, size int64) (replies []*model.ApiReply, err error) {
	replies = make([]*model.ApiReply, 0)
	rows, err := g.Mdb.Query("select cid,author_id,book_id,parent_id,root_id,commented_uid,content,create_time from comments where root_id = ? and content <>? ORDER BY stars desc limit ?,?", cid, isDel, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		reply := new(model.Comment)
		err = rows.Scan(&reply.Cid, &reply.AuthorId, &reply.BookId, &reply.ParentId, &reply.RootId, &reply.CommentedUid, &reply.Content, &reply.CreateTime)
		if err != nil {
			return nil, err
		}
		replies = append(replies, &model.ApiReply{Comment: reply})
	}
	return
}
