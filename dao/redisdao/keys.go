package redisdao

import "strconv"

func GetVerificationKey(email string) string {
	return email + ":verification"
}

func GetBookDetailKey(pid int64) string {
	return "post:" + strconv.FormatInt(pid, 10) + ":detail"
}

func GetStarBookKey(pid int64) string {
	return "post:" + strconv.FormatInt(pid, 10) + ":stars"
}

func GetStarCommentKey(cid int64) string {
	return "comment:" + strconv.FormatInt(cid, 10) + ":stars"
}
