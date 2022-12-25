package services

import (
	"errors"
	"redrock-test/business/dao"
)

func Enroll(username, password string) error {
	if dao.CheckUserExist9(username) {
		return errors.New("user has existed")
	}
	return dao.AddUser(username, password)
}
