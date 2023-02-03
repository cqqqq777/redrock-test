package redisdao

import (
	"context"
	"github.com/redis/go-redis/v9"
	g "redrock-test/global"
)

func GetCommentsStars(cid int64) (int64, error) {
	intCmd := g.Rdb.BitCount(context.Background(), GetStarCommentKey(cid), &redis.BitCount{})
	return intCmd.Val(), intCmd.Err()
}

func GetUserStarCommentStatus(cid, id int64) (int64, error) {
	intCmd := g.Rdb.GetBit(context.Background(), GetStarCommentKey(cid), id)
	return intCmd.Val(), intCmd.Err()
}
