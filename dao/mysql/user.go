package mysql

import (
	g "redrock-test/global"
	"redrock-test/model"
)

// CheckUsername 检查用户是否存在
func CheckUsername(username string) (err error) {
	var count int8
	if err = g.Mdb.QueryRow("select count(uid) from users where username = ?", username).Scan(&count); err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// CheckEmail 检查邮箱是否注册
func CheckEmail(email string) (err error) {
	var count int8
	if err = g.Mdb.QueryRow("select count(uid) from users where email = ?", email).Scan(&count); err != nil {
		return
	}
	if count > 0 {
		return ErrorEmailExist
	}
	return
}

// AddUser 新增用户
func AddUser(user *model.User) (err error) {
	_, err = g.Mdb.Exec("insert into users(uid,username,password,email) values(?,?,?,?)", user.Uid, user.Username, user.Password, user.Email)
	return
}

// FindPasswordByEmail 通过邮箱找密码
func FindPasswordByEmail(email string) (password string, err error) {
	err = g.Mdb.QueryRow("select password from users where email = ?", email).Scan(&password)
	return
}

// FindPasswordByUsername 通过用户名找密码
func FindPasswordByUsername(username string) (password string, err error) {
	err = g.Mdb.QueryRow("select password from users where username = ?", username).Scan(&password)
	return
}

// FindPasswordByUid 通过uid找密码
func FindPasswordByUid(uid int) (password string, err error) {
	err = g.Mdb.QueryRow("select password from users where uid = ?", uid).Scan(&password)
	return
}

// FindUid 查询uid
func FindUid(UsernameOrEmail string) (uid int, err error) {
	err = g.Mdb.QueryRow("select uid from users where username = ? or email = ?", UsernameOrEmail, UsernameOrEmail).Scan(&uid)
	return
}

// FindUidByEmail 通过邮箱查询uid
func FindUidByEmail(email string) (uid int, err error) {
	err = g.Mdb.QueryRow("select uid from users where email = ? ", email).Scan(&uid)
	return
}
