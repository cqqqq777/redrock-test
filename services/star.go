package services

import (
	"redrock-test/dao/mysql"
	"redrock-test/dao/redisdao"
)

func StarComment(cid, uid int64) error {
	id, err := mysql.GetIdByUid(uid)
	if err != nil {
		return err
	}
	status, err := redisdao.GetUserStarCommentStatus(cid, id)
	if err != nil {
		return err
	}
	switch status {
	case 0:
		err = redisdao.StarComment(cid, id)
	case 1:
		err = redisdao.CancelStarComment(cid, id)
	}
	return err
}
