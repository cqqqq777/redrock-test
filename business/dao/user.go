package dao

import g "redrock-test/business/global"

func CheckUserExist9(username string) bool {
	var s string
	row := g.Db.QueryRow("select username from users where username = ?", username)
	if err := row.Scan(&s); err != nil {
		return false
	}
	return true
}

func AddUser(username, password string) error {
	_, err := g.Db.Exec("insert into users(username,password) values(?,?)", username, password)
	return err
}
