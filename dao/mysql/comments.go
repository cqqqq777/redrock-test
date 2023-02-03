package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

const (
	isDel = "已删除"
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
