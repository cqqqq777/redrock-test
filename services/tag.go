package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/model"
)

func TagList() ([]*model.Tag, error) {
	return mysql.TagList()
}
