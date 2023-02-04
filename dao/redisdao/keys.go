package redisdao

import "strconv"

func GetVerificationKey(email string) string {
	return email + ":verification"
}

func GetBookDetailKey(bid int64) string {
	return "post:" + strconv.FormatInt(bid, 10) + ":detail"
}

func GetStarBookKey(bid int64) string {
	return "book:" + strconv.FormatInt(bid, 10) + ":stars"
}

func GetStarCommentKey(cid int64) string {
	return "comment:" + strconv.FormatInt(cid, 10) + ":stars"
}
