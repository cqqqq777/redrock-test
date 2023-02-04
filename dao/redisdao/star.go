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

func StarComment(cid, uid int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarCommentKey(cid), uid, 1).Err()
	return err
}

func CancelStarComment(cid, uid int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarCommentKey(cid), uid, 0).Err()
	return err
}

func GetUserStarBookStatus(bid, id int64) (int64, error) {
	intCmd := g.Rdb.GetBit(context.Background(), GetStarBookKey(bid), id)
	return intCmd.Val(), intCmd.Err()
}

func StarBook(bid, id int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarBookKey(bid), id, 1).Err()
	return err
}

func CancelStarBook(bid, id int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarBookKey(bid), id, 0).Err()
	return err
}

func GetBookStars(bid int64) (int64, error) {
	intCmd := g.Rdb.BitCount(context.Background(), GetStarBookKey(bid), &redis.BitCount{})
	return intCmd.Val(), intCmd.Err()
}
