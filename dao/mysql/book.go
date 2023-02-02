package mysql

import g "redrock-test/global"

func GetAuthorIdByBid(bid int64) (uid int64, err error) {
	err = g.Mdb.QueryRow("select author_id from books where bid = ?", bid).Scan(&uid)
	return
}
