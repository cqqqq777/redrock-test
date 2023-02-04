package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"redrock-test/dao/mysql"
	"redrock-test/dao/redisdao"
	g "redrock-test/global"
)

func Cron() {
	c := cron.New()
	_, err := c.AddFunc("@every 2h", RedisToMysql)
	if err != nil {
		panic(err)
	}
	_, err = c.AddFunc("@every 2h", SyncBookScore)
	if err != nil {
		panic(err)
	}
	c.Start()
}

func RedisToMysql() {
	cidList, err := mysql.GetCidList()
	if err != nil {
		g.Logger.Warn(fmt.Sprintf("failed to sync comments stars  err:%v", err))
		return
	}
	for _, v := range cidList {
		stars, err := redisdao.GetCommentsStars(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync cid:%d stars failed err:%v", v, err))
			continue
		}
		err = mysql.SyncCommentStars(v, stars)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync cid:%d stars failed err:%v", v, err))
		}
	}
}

func SyncBookScore() {
	bidList, err := mysql.GetBidList()
	if err != nil {
		g.Logger.Warn(fmt.Sprintf("failed to sync books score  err:%v", err))
		return
	}
	for _, v := range bidList {
		stars, err := redisdao.GetBookStars(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync bid:%d score failed err:%v", v, err))
			continue
		}
		replyNum, err := mysql.GetBookCommentsTotalNum(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync bid:%d score failed err:%v", v, err))
			continue
		}
		collectNum, err := mysql.GetBookCollectNum(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync bid:%d score failed err:%v", v, err))
			continue
		}
		score := stars*283 + replyNum*22 + collectNum*333
		err = mysql.SyncBookScore(v, score)
	}
}
