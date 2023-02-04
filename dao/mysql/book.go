package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

func GetAuthorIdByBid(bid int64) (uid int64, err error) {
	err = g.Mdb.QueryRow("select author_id from books where bid = ?", bid).Scan(&uid)
	return
}

func TagBookList(tid, page, size int64) (data []*model.Book, err error) {
	data = make([]*model.Book, 0)
	rows, err := g.Mdb.Query("select bid,name,score from books where bid in (select bid from book_tag_map where tid = ? ) order by score desc limit ?,?", tid, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			bid   int64
			name  string
			score int64
		)
		err = rows.Scan(&bid, &name, &score)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.Book{BookId: bid, Name: name, Score: score})
	}
	return
}

func BookDetail(bid int64) (book *model.Book, err error) {
	book = new(model.Book)
	err = g.Mdb.QueryRow("select bid,name,author,score,cover,link,publish_time from books where bid = ?", bid).Scan(&book.BookId, &book.Name, &book.Author, &book.Score, &book.Cover, &book.Link, &book.PublishTime)
	return
}

func GetUserStarBook(bid, uid int64) (started int8, err error) {
	err = g.Mdb.QueryRow("select count(id) from bookshelf where bid=? and uid = ?", bid, uid).Scan(&started)
	return
}

func GetBookTags(bid int64) (tags string, err error) {
	rows, err := g.Mdb.Query("select tag from tags where tid in (select tid from book_tag_map where bid = ?)", bid)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return "", err
		}
		tags += tag + " "
	}
	return
}

func CollectBook(bid, uid int64) (err error) {
	_, err = g.Mdb.Exec("insert into bookshelf (uid, bid) values (?,?);", uid, bid)
	return err
}

func CancelCollectBook(bid, uid int64) error {
	_, err := g.Mdb.Exec("delete from bookshelf where bid = ? and uid = ?", bid, uid)
	return err
}

func SearchBooks(page, size int64, key string) (data []*model.Book, err error) {
	data = make([]*model.Book, 0)
	rows, err := g.Mdb.Query("select bid,name,score from books where name like ? order by score desc limit ?,?", "%"+key+"%", (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			bid   int64
			name  string
			score int64
		)
		err = rows.Scan(&bid, &name, &score)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.Book{BookId: bid, Name: name, Score: score})
	}
	return
}

func BookList(page, size int64) (data []*model.Book, err error) {
	data = make([]*model.Book, 0)
	rows, err := g.Mdb.Query("select bid,name,score from books order by score desc limit ?,?", (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			bid   int64
			name  string
			score int64
		)
		err = rows.Scan(&bid, &name, &score)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.Book{BookId: bid, Name: name, Score: score})
	}
	return
}

func GetBidList() (list []int64, err error) {
	list = make([]int64, 0)
	rows, err := g.Mdb.Query("select bid from books ")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var bid int64
		err = rows.Scan(&bid)
		if err != nil {
			return nil, err
		}
		list = append(list, bid)
	}
	return
}

func GetBookCollectNum(bid int64) (num int64, err error) {
	err = g.Mdb.QueryRow("select count(id) from bookshelf where bid = ?", bid).Scan(&num)
	return
}

func SyncBookScore(bid, score int64) error {
	_, err := g.Mdb.Exec("update books set score = ? where bid = ?", score, bid)
	return err
}
