package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

func CommentBook(comment *model.Comment) (err error) {
	_, err = g.Mdb.Exec("insert into comments (cid, author_id, book_id, commented_uid, content) values (?,?,?,?,?)", comment.Cid, comment.AuthorId, comment.BookId, comment.CommentedUid, comment.Content)
	return
}
